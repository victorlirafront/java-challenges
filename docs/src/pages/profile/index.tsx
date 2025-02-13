import ProfileContainer from '@/components/ProfileContainer/ProfileContainer';
import PrivateRoute from '@/components/PrivateRoute/PrivateRoute';

function Profile() {

  return (
    <PrivateRoute>
      <ProfileContainer />
    </PrivateRoute>
  );
}

export default Profile;
