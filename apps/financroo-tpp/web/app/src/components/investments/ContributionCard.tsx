import React from "react";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";

const useStyles = makeStyles((theme) => ({
  card: {
    width: 500,
    padding: "24px 0",
    margin: "auto",
  },
  title: {
    fontWeight: 600,
    fontSize: 16,
    lineHeight: "24px",
    color: "#212533",
    textAlign: "center",
    paddingBottom: 17,
    borderBottom: "solid 1px #ECECEC",
  },
  content: {
    padding: 32,
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  },
  buttons: {
    borderTop: "solid 1px #ECECEC",
    padding: "24px 32px 0 32px",
    display: "flex",
    justifyContent: "space-between",
    "& > button": {
      minWidth: 84,
      ...theme.custom.button,
    },
  },
  backButton: {
    color: `${theme.palette.secondary.main} !important`,
    "&:hover": {
      backgroundColor: "white",
    },
  },
  nextButton: {
    "&:hover": {
      backgroundColor: theme.palette.primary.main,
    },
  },
}));

type Props = {
  title?: string | React.ReactNode;
  children: React.ReactNode;
  backButton?: { title: string; onClick: () => void };
  nextButton?: { title: string; onClick: () => void; disabled?: boolean };
};

export default function ContributionCard({
  title,
  children,
  backButton,
  nextButton,
}: Props) {
  const classes = useStyles();

  return (
    <Card className={classes.card}>
      {title && <div className={classes.title}>{title}</div>}
      <CardContent className={classes.content}>{children}</CardContent>
      {backButton && nextButton && (
        <CardActions className={classes.buttons}>
          <Button
            onClick={backButton.onClick}
            variant="outlined"
            color="secondary"
            className={classes.backButton}
          >
            {backButton.title}
          </Button>
          <Button
            onClick={nextButton.onClick}
            variant="contained"
            disableElevation
            color="primary"
            style={{ color: "white" }}
            className={classes.nextButton}
            disabled={nextButton.disabled || false}
          >
            {nextButton.title}
          </Button>
        </CardActions>
      )}
    </Card>
  );
}
