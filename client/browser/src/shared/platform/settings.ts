import { applyEdits, modify, parse as parseJSONC } from 'jsonc-parser'
import { from, fromEvent, Observable } from 'rxjs'
import { distinctUntilChanged, filter, map, startWith } from 'rxjs/operators'

import { isErrorLike } from '@sourcegraph/common'
import { dataOrThrowErrors, gql } from '@sourcegraph/http-client'
import { SettingsEdit } from '@sourcegraph/shared/src/api/client/services/settings'
import { PlatformContext } from '@sourcegraph/shared/src/platform/context'
import {
    mergeSettings,
    SettingsCascade,
    SettingsCascadeOrError,
    SettingsSubject,
    SubjectSettingsContents,
} from '@sourcegraph/shared/src/settings/settings'

import { observeStorageKey, storage } from '../../browser-extension/web-extension-api/storage'
import { ViewerConfigurationResult } from '../../graphql-operations'
import { isInPage } from '../context'

const inPageClientSettingsKey = 'sourcegraphClientSettings'

/**
 * Returns an observable that emits the localStorage value (as a raw string) for the given key on
 * every storage update event, starting with the current value.
 */
function observeLocalStorageKey(key: string, defaultValue: string): Observable<string> {
    const initialValue = localStorage.getItem(key)
    return fromEvent<StorageEvent>(window, 'storage').pipe(
        filter(event => event.key === key),
        map(event => event.newValue),
        startWith(initialValue),
        map(value => value ?? defaultValue),
        distinctUntilChanged()
    )
}

const createStorageSettingsCascade: () => Observable<SettingsCascade> = () => {
    /** Observable of the JSONC string of the settings.
     *
     * NOTE: We can't use LocalStorageSubject here because the JSONC string is stored raw in localStorage and LocalStorageSubject also does parsing.
     * This could be changed, but users already have settings stored, so it would need a migration for little benefit.
     */
    const storageObservable = isInPage
        ? observeLocalStorageKey(inPageClientSettingsKey, '{}')
        : observeStorageKey('sync', 'clientSettings')

    const subject: SettingsSubject = {
        __typename: 'Client',
        id: 'Client',
        displayName: 'Client',
        viewerCanAdminister: true,
    }

    return storageObservable.pipe(
        map(clientSettingsString => parseJSONC(clientSettingsString || '')),
        map(clientSettings => ({
            subjects: [
                {
                    subject,
                    settings: clientSettings,
                    lastID: null,
                },
            ],
            final: clientSettings || {},
        }))
    )
}

/**
 * The settings cascade consisting solely of client settings.
 */
export const storageSettingsCascade = createStorageSettingsCascade()

/**
 * Merge two settings cascades (used to merge viewer settings and client settings).
 */
export function mergeCascades(
    cascadeOrError: SettingsCascadeOrError,
    cascade: SettingsCascade
): SettingsCascadeOrError {
    return {
        subjects:
            cascadeOrError.subjects === null
                ? cascade.subjects
                : isErrorLike(cascadeOrError.subjects)
                ? cascadeOrError.subjects
                : [...cascadeOrError.subjects, ...cascade.subjects],
        final:
            cascadeOrError.final === null
                ? cascade.final
                : isErrorLike(cascadeOrError.final)
                ? cascadeOrError.final
                : mergeSettings([cascadeOrError.final, cascade.final]),
    }
}

// This is a fragment on the DEPRECATED GraphQL API type ConfigurationCascade (not SettingsCascade) for backcompat.
const configurationCascadeFragment = gql`
    fragment ConfigurationCascadeFields on ConfigurationCascade {
        subjects {
            __typename
            ... on Org {
                name
                displayName
            }
            ... on User {
                username
                displayName
            }
            ... on Site {
                siteID
                allowSiteSettingsEdits
            }
            id
            latestSettings {
                id
                contents
            }
            settingsURL
            viewerCanAdminister
        }
        merged {
            contents
            messages
        }
    }
`

/**
 * Fetches the settings cascade for the viewer.
 *
 * TODO(sqs): This uses the DEPRECATED GraphQL Query.viewerConfiguration and ConfigurationCascade for backcompat.
 */
export function fetchViewerSettings(requestGraphQL: PlatformContext['requestGraphQL']): Observable<{
    final: string
    subjects: (SettingsSubject & SubjectSettingsContents)[]
}> {
    return from(
        requestGraphQL<ViewerConfigurationResult>({
            request: gql`
                query ViewerConfiguration {
                    viewerConfiguration {
                        ...ConfigurationCascadeFields
                    }
                }
                ${configurationCascadeFragment}
            `,
            variables: {},
            mightContainPrivateInfo: false,
        })
    ).pipe(
        map(dataOrThrowErrors),
        map(({ viewerConfiguration }) => {
            if (!viewerConfiguration) {
                throw new Error('fetchViewerSettings: empty viewerConfiguration')
            }

            for (const subject of viewerConfiguration.subjects) {
                // User/org/global settings cannot be edited from the
                // browser extension (only client settings can).
                subject.viewerCanAdminister = false
            }

            return {
                subjects: viewerConfiguration.subjects,
                final: viewerConfiguration.merged.contents,
            }
        })
    )
}

/**
 * Applies an edit and persists the result to client settings.
 */
export async function editClientSettings(edit: SettingsEdit | string): Promise<void> {
    const getNext = (previous: string): string =>
        typeof edit === 'string'
            ? edit
            : applyEdits(
                  previous,
                  modify(previous, edit.path, edit.value, {
                      formattingOptions: {
                          tabSize: 2,
                          insertSpaces: true,
                          eol: '\n',
                      },
                  })
              )
    if (isInPage) {
        const previous = localStorage.getItem(inPageClientSettingsKey) || ''
        const next = getNext(previous)

        localStorage.setItem(inPageClientSettingsKey, next)

        return Promise.resolve()
    }

    const { clientSettings: previous = '{}' } = await storage.sync.get()
    const next = getNext(previous)

    await storage.sync.set({ clientSettings: next })
}
