import React from "react";
import logo from "./logo.svg";
import "./App.css";

import ApolloClient from "apollo-boost";
import { ApolloProvider } from "@apollo/react-hooks";
import Test from "./Test";

function App() {
  const client = new ApolloClient({
    uri: "http://localhost:4242/query",
  });
  return (
    <ApolloProvider client={client}>
      <div>
        <Test />
      </div>
    </ApolloProvider>
  );
}

export default App;