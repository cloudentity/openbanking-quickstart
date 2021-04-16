import React from "react";
import Toolbar from "@material-ui/core/Toolbar";
import AppBar from "@material-ui/core/AppBar";
import logo from '../assets/cloudentity-white.svg';

export default function PageToolbar({children}) {

  return (
    <AppBar position="fixed" variant={'outlined'}>
      <Toolbar>
        <div>
          <img src={logo} alt={'logo'} style={{width: 160}}/>
        </div>
        {children}
      </Toolbar>
    </AppBar>
  )
};
