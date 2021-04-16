import React from "react";
import {Typography} from "@material-ui/core";

type Props = {
  id: string
  label: string,
  caption?: string,
  icon?: any
}

export default function FormInputLabel({id, label, icon: Icon, caption}: Props) {
  return (
    <div id={id} style={{fontSize: 12, fontWeight: 600, marginBottom: 12, color: '#212533', position: 'relative', display: 'flex', justifyContent: 'space-between'}}>
      <div>{label} {Icon && <Icon style={{color: '#4CAF50', width: 16, height: 16, marginLeft: 8, position: 'absolute'}}/>}</div>
      {caption && <Typography variant={'caption'}>{caption}</Typography>}
    </div>
  )
};
