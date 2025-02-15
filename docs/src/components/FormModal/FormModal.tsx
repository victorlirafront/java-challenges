import { StyledFormModal } from './FormModal.styled';
import Image from 'next/image';
import { FormModalProps } from './FormModal.types'

function FormModal(props: FormModalProps) {

  const { className } = props

  return (
    <StyledFormModal className={className}>
      <Image
        src="https://ik.imagekit.io/Victorliradev/blog-api-golang/assets/check_pi6qcB_fQ.svg?updatedAt=1739627820697"
        alt="Ícone de confirmação"
        width={30}
        height={30}
      />
      <p>Você foi autenticado com sucesso!</p>
    </StyledFormModal>
  );
}

export default FormModal;
