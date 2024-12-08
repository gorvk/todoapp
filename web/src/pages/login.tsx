import { useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { login, callback } from "../svc/auth";

export function Login(props: { isCallback: boolean }) {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();

  const { isCallback } = props;

  useEffect(() => {
    const isLoggedIn = async () => {
      const res = await login();
      window.location.href = res.url;
    };

    const callbackData = async () => {
      const code = searchParams.get("code");
      if (code) {
        await callback(code);
        navigate("/home");
      }
    };

    if (isCallback) {
      callbackData();
      return;
    }

    isLoggedIn();
    // eslint-disable-next-line
  }, []);

  return <h3 className="text-white">LOADING...</h3>;
}
