
import PrivateRoute from '@/components/ProtectedRoute/ProtectedRoute';
import ProfileContainer from './Components/ProfileContainer/ProfileContainer';

function Profile() {
  return (
    <PrivateRoute>
      <ProfileContainer />
    </PrivateRoute>
  );
}

export default Profile;
