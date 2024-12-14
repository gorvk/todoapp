import { Link } from "react-router-dom";

export function Welcome() {
  return (
    <>
      <h3>Welcome to Todo app</h3>
      <ul>
        <li>
          <Link to='/dashboard' className="text-green-500">DASHBOARD</Link>
        </li>
      </ul>
    </>
  );
}
