import { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useAuth0 } from "@auth0/auth0-react";
import { slice as userSlice } from "../state/slices/user";
import { AppState } from "../state/store";

export function Auth() {
  const { loginWithPopup, getIdTokenClaims } = useAuth0();
  const dispatch = useDispatch();
  const currentUser = useSelector((state: AppState) => state.currentUser);

  useEffect(() => {
    const authenticateUser = async () => {
      await loginWithPopup();
      const claim = await getIdTokenClaims();
      if (claim) {
        const idToken = claim.__raw;
        localStorage.setItem('id_token', idToken)
        dispatch(userSlice.actions.setCurrentUser(claim));
      }
    };
    if (!currentUser) {
      authenticateUser();
    }
  }, []);

  return <h3 className="text-white">LOADING...</h3>;
}
