import copy from 'clipboard-copy';
import {showNotification} from '../components/notifications/actions/actions';
import getStore from '../store';

export const copyToClipboard = (value, label) => {
  copy(value)
    .then(() => {
      getStore().dispatch(showNotification({
        message: `${label} value copied to clipboard.`,
        options: {
          variant: 'info'
        }
      }))
    })
};
