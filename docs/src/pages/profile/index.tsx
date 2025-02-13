import Dashboard from "@/components/Dashboard/Dashboard";
import PrivateRoute from "@/components/PrivateRoute/PrivateRoute";

function Profile() {
  return (
    <PrivateRoute>
      <Dashboard />
    </PrivateRoute>
  );
}

export default Profile;
