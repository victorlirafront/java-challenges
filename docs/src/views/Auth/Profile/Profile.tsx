
import PrivateRoute from '@/components/PrivateRoute/PrivateRoute';
import ProfileContainer from './Components/ProfileContainer/ProfileContainer';

function Profile() {
  return (
    <PrivateRoute>
      <ProfileContainer />
    </PrivateRoute>
  );
}

export default Profile;
