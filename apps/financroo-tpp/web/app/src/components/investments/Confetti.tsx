import React, { useEffect, useRef } from "react";
import { makeStyles } from "tss-react/mui";
import confetti from "canvas-confetti";

const useStyles = makeStyles()(() => ({
  canvas: {
    position: "absolute",
    top: 0,
  },
}));

export default function Confetti() {
  const canvasRef = useRef(null);
  const { classes } = useStyles();

  useEffect(() => {
    if (canvasRef) {
      setTimeout(() => {
        confetti.create(canvasRef.current)({
          spread: 360,
          origin: {
            y: 0.2,
          },
          startVelocity: 10,
        });
      }, 500);
    }
  }, [canvasRef]);

  return (
    <canvas
      className={classes.canvas}
      ref={canvasRef}
      width={500}
      height={275}
    />
  );
}
