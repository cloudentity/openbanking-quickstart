import React from 'react';
import {CircularProgress, withStyles} from '@material-ui/core';

const styles = theme => ({
  progress: {
    width: 100,
    height: 100,
    position: 'absolute',
    left: 'calc(50% - 50px);',
  },
  circle: {
    position: 'absolute',
    top: 0,
    left: 0,
  },
});


const Progress = ({size = 100, top = 40, classes}) => (
  <div className={classes.progress} style={{top}}>
    <CircularProgress size={size} className={classes.circle} thickness={3}/>
  </div>
);

export default withStyles(styles, {withTheme: true})(Progress);
