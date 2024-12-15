import { createRoot } from "react-dom/client";
import App from "./App.tsx";
import { Provider } from "react-redux";
import { store } from "./state/store.ts";
import { Auth0Provider } from "@auth0/auth0-react";

createRoot(document.getElementById("root") as HTMLElement).render(
  <Auth0Provider
    domain={`${process.env.AUTH0_DOMAIN}`}
    clientId={`${process.env.AUTH0_CLIENT_ID}`}
    authorizationParams={{
      redirect_uri: window.location.origin + '/dashboard',
    }}
  >
    <Provider store={store}>
      <App />
    </Provider>
  </Auth0Provider>
);
