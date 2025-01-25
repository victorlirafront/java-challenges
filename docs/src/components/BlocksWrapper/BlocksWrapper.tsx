import { ReactNode } from "react";
import { StyledBlocksWrapper } from "./BlocksWrapper.styled";

type BlocksWrapperProps = {
  children: ReactNode
}

function BlocksWrapper(props: BlocksWrapperProps){
  return <StyledBlocksWrapper>
    {props.children}
  </StyledBlocksWrapper>
}

export default BlocksWrapper