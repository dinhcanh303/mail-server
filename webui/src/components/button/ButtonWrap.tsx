interface ButtonWrapProps {
  href?: string
  children?: React.ReactNode
}

const ButtonWrap: React.FC<ButtonWrapProps> = ({ href, children }) => {
  return (
    <div>
      {href ? (
        <a
          href={href}
          className='ani1mate-pulse flex items-center space-x-2 p-1 cursor-pointer hover:bg-gray-200 dark:hover:bg-gray-800 rounded-lg transition-all'
        >
          {children}
        </a>
      ) : (
        <div className='ani1mate-pulse flex items-center space-x-2 p-1 cursor-pointer hover:bg-gray-200 dark:hover:bg-gray-800 rounded-lg transition-all'>
          {children}
        </div>
      )}
    </div>
  )
}
export default ButtonWrap
