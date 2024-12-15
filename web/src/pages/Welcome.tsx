import { Link } from "react-router-dom";

export function Welcome() {
  return (
    <>
      <h3>Welcome to Todo app</h3>
      <ul>
        <li>
          <div>
            <Link to="/dashboard" className="text-green-500">
              DASHBOARD
            </Link>
          </div>
          <Link to="/auth">
            <button className="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-4 border border-blue-500 hover:border-transparent rounded">
              Login
            </button>
          </Link>
          <Link to="/dashboard">
            <button className="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-4 border border-blue-500 hover:border-transparent rounded">
              Dashboard
            </button>
          </Link>
        </li>
      </ul>
    </>
  );
}
