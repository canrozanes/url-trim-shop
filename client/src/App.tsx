import React from "react";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import NotFound from "./pages/404";
import Home from "./pages/home";

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/404">
          <NotFound />
        </Route>
        <Route path="/">
          <Home />
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
