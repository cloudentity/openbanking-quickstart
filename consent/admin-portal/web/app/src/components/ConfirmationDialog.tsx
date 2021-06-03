import Button from "@material-ui/core/Button";
import Dialog from "@material-ui/core/Dialog";
import DialogContent from "@material-ui/core/DialogContent";
import DialogTitle from "@material-ui/core/DialogTitle";
import DoneIcon from "@material-ui/icons/Done";
import React from "react";
import { Typography } from "@material-ui/core";

type Props = {
  title?: string;
  content?: string;
  warningItems?: string[] | undefined;
  confirmText?: string;
  cancelText?: string;
  extraActionText?: string;
  onExtraAction?: () => void;
  onCancel: () => void;
  onConfirm: () => void;
};

const ConfirmationDialog = ({
  title,
  content,
  warningItems,
  confirmText,
  cancelText,
  extraActionText,
  onExtraAction,
  onCancel,
  onConfirm,
}: Props) => {
  return (
    <Dialog
      open={true}
      onClose={onCancel}
      aria-labelledby="confirmation-dialog-title"
      aria-describedby="confirmation-dialog-description"
      maxWidth={"sm"}
      fullWidth
    >
      <DialogTitle
        disableTypography
        id="confirmation-dialog-title"
        style={{
          color: "#BD271E",
          background: "#F7FAFF",
          padding: "24px 32px",
        }}
      >
        <strong>{title}</strong>
      </DialogTitle>
      <DialogContent style={{ padding: "24px 32px" }}>
        <Typography
          id="confirmation-dialog-content"
          component={"div"}
          variant="body1"
        >
          <strong>{content}</strong>
        </Typography>
        {warningItems && (
          <div style={{ marginTop: 8 }}>
            {warningItems.map((i) => (
              <div
                style={{ marginTop: 4, display: "flex", alignItems: "center" }}
              >
                <DoneIcon style={{ marginRight: 6, color: "#BD271E" }} />{" "}
                <Typography style={{ fontSize: 14 }}>{i}</Typography>
              </div>
            ))}
          </div>
        )}
        <div style={{ marginTop: 46, display: "flex" }}>
          <Button
            id="confirm-button"
            fullWidth
            onClick={() => onConfirm()}
            color="secondary"
            style={{ background: "#BD271E" }}
            variant={"contained"}
          >
            {confirmText || "Confirm"}
          </Button>
          {onExtraAction && (
            <Button
              id="confirm-button"
              fullWidth
              onClick={() => onExtraAction()}
              color="primary"
              style={{ marginLeft: 12 }}
              variant={"contained"}
            >
              {extraActionText}
            </Button>
          )}
          <Button
            id="cancel-button"
            fullWidth
            onClick={() => onCancel()}
            color="primary"
            style={{ marginLeft: 12 }}
            variant={"outlined"}
          >
            {cancelText || "Cancel"}
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  );
};

export default ConfirmationDialog;
