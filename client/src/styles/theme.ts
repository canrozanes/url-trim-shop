import { createMuiTheme } from '@material-ui/core/styles';

const theme = createMuiTheme({
  palette: {
    primary: {
      main: '#0052CC',
      // textColor: '#172B4D',
      contrastText: '#fff',
      light: '#FFF',
      // dark: '#002884',
    },
    secondary: {
      main: "#EAC435"
    },
  },
  typography: {
    fontFamily: 'Nunito Sans, sans-serif',
    fontSize: 14
  },
});

export default theme;