import React from 'react';
import {useForm} from 'react-hook-form';
import {validators} from './validation';
import ComplexToggle from "./ComplexToggle";
import commonTextField from "./CommonTextField";
import FormFooter from "./FormFooter";
import CheckBox from "./CheckBox";
import SwitchField from "./SwitchField";
import AutocompleteField from "./AutocompleteField";
import SelectField from "./SelectField";
import SwitchBlock from "./SwitchBlock";

export const useFormFactory = ({id, data = {}}) => {
  const form = useForm({
    defaultValues: data
  });

  const createRequiredField = ({validate = {}, ...props}: any) => commonTextField({
    validate: {
      length: validators.length({label: props.label}),
      maxLength: validators.maxLength({label: props.label}),
      ...validate
    },
    id, form, ...props,
  });

  const createField = ({validate = {}, ...props}: any) => commonTextField({
    validate: {
      maxLength: validators.maxLength({label: props.label}),
      ...validate

    },
    id, form, ...props
  });

  const createSelect = props => SelectField({id, form, ...props})

  const createRequiredSelect = ({validate = {}, ...props}: any) => SelectField({
    validate: {
      required: `${props.label} is required`,
      ...validate
    },
    id, form, ...props
  })

  const createCheckBox = props => CheckBox({id, form, ...props});

  const createSwitch = props => SwitchField({id, form, ...props});

  const createSwitchBlock = props => <SwitchBlock id={id} form={form} {...props}/>

  const createComplexToggle = props => ComplexToggle({id, form, ...props});

  const createReadOnlyField = (props) => commonTextField({id, form, ...props, disabled: true});

  const createFormFooter = props => <FormFooter id={id} form={form} {...props}/>

  const createAutocompleteField = props => AutocompleteField({id, form, ...props});

  return {
    ...form,
    createRequiredField,
    createField,
    createSelect,
    createRequiredSelect,
    createCheckBox,
    createSwitch,
    createSwitchBlock,
    createAutocompleteField,
    createReadOnlyField,
    createComplexToggle,
    createFormFooter
  }
};
