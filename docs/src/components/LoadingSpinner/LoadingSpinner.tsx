import Image from "next/image"
import { StyledLoadingSpinner } from "./LoadingSpinner.styled"

function LoadingSpinner(){
  return (
    <StyledLoadingSpinner>
    <Image width={150} height={150} alt="loading spinner" src='https://ik.imagekit.io/Victorliradev/blog-api-golang/assets/loading_zNT_aRbVJ.gif?updatedAt=1739733698755' />
    </StyledLoadingSpinner>
  )
}

export default LoadingSpinner