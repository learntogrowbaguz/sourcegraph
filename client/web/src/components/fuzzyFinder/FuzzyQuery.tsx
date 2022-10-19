import { CaseInsensitiveFuzzySearch } from '../../fuzzyFinder/CaseInsensitiveFuzzySearch'
import { FuzzySearch, IndexingFSM, SearchIndexing, SearchValue } from '../../fuzzyFinder/FuzzySearch'

import { FuzzyFSM } from './FuzzyFsm'
import { FuzzyLocalCache, PersistableQueryResult } from './FuzzyLocalCache'

export abstract class FuzzyQuery {
    protected queries: Map<string, Promise<PersistableQueryResult[]>> = new Map()
    protected doneQueries: Set<string> = new Set()
    protected queryResults: Map<string, PersistableQueryResult> = new Map()
    constructor(private readonly onNamesChanged: () => void, private readonly cache: FuzzyLocalCache) {
        this.addQuery('FuzzyQuery.fromCache()-constructor', this.cache.initialValues())
    }

    protected abstract searchValues(): SearchValue[]
    protected abstract rawQuery(userQuery: string): string
    protected abstract handleRawQueryPromise(query: string): Promise<PersistableQueryResult[]>

    public async removeStaleResults(): Promise<void> {
        const fromCache = await this.cache.initialValues()
        if (fromCache.length === 0) {
            // Nothing to invalidate.
            return
        }
        const toRemove = await this.cache.staleValues(fromCache)
        this.removeQueryResults(toRemove)
    }
    private removeQueryResults(toRemove: PersistableQueryResult[]): void {
        const oldSize = this.queryResults.size
        for (const result of toRemove) {
            this.queryResults.delete(result.url)
        }
        const didChange = this.queryResults.size < oldSize
        if (didChange) {
            this.cache.cacheValues([...this.queryResults.values()])
            this.onNamesChanged()
        }
    }
    private addQueryResults(results: PersistableQueryResult[]): void {
        const oldSize = this.queryResults.size
        for (const result of results) {
            this.queryResults.set(result.url, result)
        }
        const didChangeSize = this.queryResults.size > oldSize
        if (didChangeSize) {
            this.cache.cacheValues([...this.queryResults.values()])
        }
    }
    public isDoneDownloading(): boolean {
        return this.queries.size === 0
    }
    public hasQuery(query: string): boolean {
        return this.queries.has(query) || this.doneQueries.has(query)
    }
    public fuzzySearch(): FuzzySearch {
        return new CaseInsensitiveFuzzySearch(this.searchValues(), undefined)
    }

    public fuzzyFSM(query: string): FuzzyFSM {
        this.handleQuery(query)
        if (this.queries.size === 0) {
            return {
                key: 'ready',
                fuzzy: this.fuzzySearch(),
            }
        }
        return {
            key: 'indexing',
            indexing: this.indexingFSM(),
        }
    }

    public indexingFSM(): SearchIndexing {
        let indexingPromise: Promise<IndexingFSM> | undefined
        return {
            key: 'indexing',
            partialFuzzy: this.fuzzySearch(),
            indexedFileCount: this.queryResults.size,
            totalFileCount: this.queryResults.size + 10,
            isIndexing: () => indexingPromise !== undefined,
            continueIndexing: () => {
                if (!indexingPromise) {
                    indexingPromise = Promise.any([...this.queries.values()]).then(
                        () =>
                            this.isDoneDownloading() ? { key: 'ready', value: this.fuzzySearch() } : this.indexingFSM(),
                        () => this.indexingFSM()
                    )
                }
                return indexingPromise
            },
        }
    }

    public handleQuery(query: string): void {
        if (query === '') {
            return
        }
        const actualQuery = this.rawQuery(query)
        if (this.hasQuery(actualQuery)) {
            return
        }
        this.addQuery(actualQuery, this.handleRawQueryPromise(actualQuery))
    }

    public addQuery(query: string, promise: Promise<PersistableQueryResult[]>): void {
        this.queries.set(query, promise)
        promise
            .then(
                result => this.addQueryResults(result),
                // eslint-disable-next-line no-console
                error => console.error(`failed to download results for query ${query}`, error)
            )
            .finally(() => {
                this.queries.delete(query)
                this.doneQueries.add(query)
                this.onNamesChanged()
            })
        return
    }
}
