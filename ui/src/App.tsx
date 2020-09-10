import React from "react";

import ApolloClient from "apollo-boost";
import { ApolloProvider } from "@apollo/react-hooks";
import Test from "./Test";
import { ThemeProvider } from "theme-ui";
import { Box } from "rebass";
import theme from "./theme";
import RecipeList from "./pages/RecipeList";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import RecipeDetail from "./pages/RecipeDetail";
import NavBar from "./components/NavBar";
import IngredientList from "./pages/IngredientList";
import CreateRecipe from "./pages/CreateRecipe";
import Food from "./pages/Food";

import "./tailwind.output.css";

function App() {
  const client = new ApolloClient({
    uri: process.env.REACT_APP_GQL_URL,
  });
  return (
    <ThemeProvider theme={theme}>
      <ApolloProvider client={client}>
        <Router>
          <NavBar />
          <Box
            sx={{
              maxWidth: "80%",
              mx: "auto",
              sx: 3,
            }}
          >
            <Switch>
              <Route path="/recipe/:uuid">
                <RecipeDetail />
              </Route>
              <Route path="/recipes">
                <RecipeList />
              </Route>
              <Route path="/ingredients">
                <IngredientList />
              </Route>
              <Route path="/create">
                <CreateRecipe />
              </Route>
              <Route path="/food">
                <Food />
              </Route>
              <Route path="/">
                <Test />
              </Route>
            </Switch>
          </Box>
        </Router>

        <hr />
      </ApolloProvider>
    </ThemeProvider>
  );
}

export default App;
