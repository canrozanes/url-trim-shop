import { createMuiTheme } from '@material-ui/core/styles';

const theme = createMuiTheme({
  palette: {
    primary: {
      main: '#0052CC',
      // textColor: '#172B4D',
      contrastText: '#fff',
      // light: '#',
      // dark: '#002884',
    },
    // secondary: {
    // },
  },
  typography: {
    fontFamily: 'Nunito Sans, sans-serif',
  },
});

export default theme;