import React from "react";

import ApolloClient from "apollo-boost";
import { ApolloProvider } from "@apollo/react-hooks";
import Test from "./Test";
import RecipeList from "./pages/RecipeList";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import RecipeDetail from "./pages/RecipeDetail";
import NavBar from "./components/NavBar";
import IngredientList from "./pages/IngredientList";
import CreateRecipe from "./pages/CreateRecipe";
import Food from "./pages/Food";
import Playground from "./pages/Playground";

import "./tailwind.output.css";
import { RestfulProvider } from "restful-react";
import Photos from "./pages/Photos";
import Meals from "./pages/Meals";
import { getAPIURL, getGQLURL } from "./config";
import { CookiesProvider } from "react-cookie";

function App() {
  const client = new ApolloClient({
    uri: getGQLURL(),
  });
  return (
    <CookiesProvider>
      <RestfulProvider base={getAPIURL()}>
        <ApolloProvider client={client}>
          <Router>
            <NavBar />
            <div className="container mx-auto">
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
                <Route path="/playground">
                  <Playground />
                </Route>
                <Route path="/photos">
                  <Photos />
                </Route>
                <Route path="/meals">
                  <Meals />
                </Route>
                <Route path="/">
                  <Test />
                </Route>
              </Switch>
            </div>
          </Router>

          <hr />
        </ApolloProvider>
      </RestfulProvider>
    </CookiesProvider>
  );
}

export default App;
