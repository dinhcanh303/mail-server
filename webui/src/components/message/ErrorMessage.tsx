interface ErrorMessageProps {
  message?: string
}
const ErrorMessage: React.FC<ErrorMessageProps> = ({ message = '' }) => {
  return <div className='text-red-500 text-sm mt-1 ml-1'>{message}</div>
}
export default ErrorMessage
