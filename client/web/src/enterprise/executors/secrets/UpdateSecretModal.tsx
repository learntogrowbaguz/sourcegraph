import React, { useCallback, useState } from 'react'

import { ErrorAlert } from '@sourcegraph/branded/src/components/alerts'
import { Form } from '@sourcegraph/branded/src/components/Form'
import { logger } from '@sourcegraph/common'
import { Button, Modal, Input } from '@sourcegraph/wildcard'

import { LoaderButton } from '../../../components/LoaderButton'
import { ExecutorSecretFields } from '../../../graphql-operations'

import { useUpdateExecutorSecret } from './backend'
import { ModalHeader } from './ModalHeader'

interface UpdateSecretModalProps {
    secret: ExecutorSecretFields

    onCancel: () => void
    afterUpdate: () => void
}

export const UpdateSecretModal: React.FunctionComponent<React.PropsWithChildren<UpdateSecretModalProps>> = ({
    secret,
    onCancel,
    afterUpdate,
}) => {
    const labelId = 'updateSecret'

    const [value, setValue] = useState<string>('')
    const onChangeValue = useCallback<React.ChangeEventHandler<HTMLInputElement>>(event => {
        setValue(event.target.value)
    }, [])

    const [updateExecutorSecret, { loading, error }] = useUpdateExecutorSecret()

    const onSubmit = useCallback<React.FormEventHandler>(
        async event => {
            event.preventDefault()

            try {
                await updateExecutorSecret({
                    variables: {
                        secret: secret.id,
                        scope: secret.scope,
                        value,
                    },
                })

                afterUpdate()
            } catch (error) {
                logger.error(error)
            }
        },
        [updateExecutorSecret, secret.id, secret.scope, value, afterUpdate]
    )
    return (
        <Modal onDismiss={onCancel} aria-labelledby={labelId}>
            <ModalHeader id={labelId} secretKey={secret.key} />

            {error && <ErrorAlert error={error} />}

            <Form onSubmit={onSubmit}>
                <div className="form-group">
                    <Input
                        id="value"
                        name="value"
                        type="password"
                        autoComplete="off"
                        required={true}
                        spellCheck="false"
                        minLength={1}
                        label="Value"
                        placeholder="******"
                        value={value}
                        onChange={onChangeValue}
                    />
                </div>
                <div className="d-flex justify-content-end">
                    <Button disabled={loading} className="mr-2" onClick={onCancel} outline={true} variant="secondary">
                        Cancel
                    </Button>
                    <LoaderButton
                        type="submit"
                        disabled={loading || value.length === 0}
                        variant="primary"
                        loading={loading}
                        alwaysShowLabel={true}
                        label="Update secret"
                    />
                </div>
            </Form>
        </Modal>
    )
}
