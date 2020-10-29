import React from "react"
import { BrowserRouter as Router, Switch, Route } from "react-router-dom"
import CssBaseline from "@material-ui/core/CssBaseline"
import MainLayout from "./layouts/main"
import { ThemeProvider } from "@material-ui/core/styles"
import theme from "./styles/theme"
import NotFound from "./pages/404"
import Home from "./pages/home"

function App() {
  return (
    <>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <Router>
          <MainLayout>
            <Switch>
              <Route exact path="/">
                <Home />
              </Route>
              <Route>
                <NotFound />
              </Route>
            </Switch>
          </MainLayout>
        </Router>
      </ThemeProvider>
    </>
  )
}

export default App
