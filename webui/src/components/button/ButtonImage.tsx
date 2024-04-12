import { Button } from 'primereact/button'

interface ButtonImageProps {
  label?: string
  className?: string
  severity?: 'secondary' | 'success' | 'info' | 'warning' | 'danger' | 'help'
  srcImage: string
  height?: number
  width?: number
  alt?: string
}
const ButtonImage: React.FC<ButtonImageProps> = ({
  label,
  className = 'rounded-lg mx-2',
  severity = 'secondary',
  srcImage,
  height = 20,
  width = 20,
  alt = 'image'
}) => {
  return (
    <Button
      className={className}
      label={label}
      severity={severity}
      outlined
      icon={<img className='mr-2' src={srcImage} height={height} width={width} alt={alt} />}
    />
  )
}

export default ButtonImage
