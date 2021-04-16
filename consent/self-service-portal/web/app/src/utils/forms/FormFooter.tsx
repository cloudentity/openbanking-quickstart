import Button from "@material-ui/core/Button";
import React from "react";

interface CreateFormFooterProps {
  id: string
  form: any
  cancelText?: string
  submitText?: string
  submitButtonColor?: 'primary' | 'secondary'
  disabled?: boolean
  onCancel?: () => void
  onSubmit?: (data: any) => void
  align?: 'left' | 'center' | 'right'
  buttonWidth?: number
  style?: any
}

export default function FormFooter({id, form, cancelText, submitText, submitButtonColor = 'primary', disabled, onCancel, onSubmit, align = 'left', buttonWidth = 192, style = {}}: CreateFormFooterProps) {
  return (
    <div style={{textAlign: align, ...style}}>
      {onCancel && (
        <Button id={`${id}-cancel-button`}
                variant={'outlined'}
                color='primary'
                size="large"
                onClick={onCancel}
                style={{width: buttonWidth, marginRight: 14}}>{cancelText || 'Cancel'}</Button>
      )}
      {onSubmit && (
        <Button id={`${id}-confirm-button`}
                variant={'contained'}
                color={submitButtonColor}
                size="large"
                disabled={disabled}
                onClick={form.handleSubmit(onSubmit)}
                style={{width: buttonWidth}}>{submitText || 'Next'}</Button>
      )}
    </div>
  );
}
