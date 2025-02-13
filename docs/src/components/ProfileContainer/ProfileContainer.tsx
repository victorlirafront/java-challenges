import React from "react";

import { useRouter } from "next/router";
import { useAuth } from "@/context/Auth";
import StyledProfileContainer from "./ProfileContainer.styled";

function ProfileContainer() {
  const { logout } = useAuth();
  const router = useRouter();

  const handleLogout = () => {
    logout();
    router.push("/login");
  };

  return (
    <StyledProfileContainer>
      <h1>Dashboard</h1>
      <button onClick={handleLogout}>Logout</button>
    </StyledProfileContainer>
  );
}

export default ProfileContainer;