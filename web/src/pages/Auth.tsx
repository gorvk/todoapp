import { useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { callback, login } from "../svc/auth";
import { useDispatch } from "react-redux";
import { slice as userSlice } from "../state/slices/user";

export function Auth(props: { isCallback: boolean }) {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const dispatch = useDispatch();
  const { isCallback } = props;

  useEffect(() => {
    const callbackData = async () => {
      const code = searchParams.get("code");
      if (code) {
        const profile = await callback(code);
        dispatch(userSlice.actions.setCurrentUser(profile));
        navigate("/dashboard");
      }
    };

    if (isCallback) {
      callbackData();
      return;
    } else {
      (async () => {
        const res = await login();
        window.location.href = res.url;
      })();
    }
  }, []);

  return <h3 className="text-white">LOADING...</h3>;
}
