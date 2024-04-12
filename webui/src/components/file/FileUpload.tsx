import { upload } from '@/apis/upload.api'
import { ApiResponse } from '@/types/app'
import { isAxiosError } from 'axios'
import { Button } from 'primereact/button'
import { useRef, useState } from 'react'
import toast from 'react-hot-toast'

interface FileUploadProps {
  mode?: 'basic' | 'advance'
  icon?: string
  rounded?: boolean
  multiple?: boolean
  size?: 'large' | 'small'
  label?: string
  onUploaded?: any
  onOpenWriteComment?: any
  className?: string
}
const FileUpload: React.FC<FileUploadProps> = ({
  mode = 'basic',
  icon,
  rounded = false,
  size = 'small',
  label,
  multiple = false,
  onUploaded,
  onOpenWriteComment,
  className
}) => {
  const fileInputRef = useRef<HTMLInputElement | null>(null)
  const handleButtonClick = () => {
    // Trigger file input click when the button is clicked
    if (fileInputRef.current) {
      fileInputRef.current.click()
    }
  }
  const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      const selectedFile = e.target.files[0]
      const formData = new FormData()
      if (selectedFile) {
        formData.append('files', selectedFile)
        try {
          const res = await upload(formData)
          if (res?.status == 200) {
            const attachments = res.data.map((item, _) => {
              return item
            })
            if (onOpenWriteComment) {
              onOpenWriteComment(true)
            }
            onUploaded(attachments)
            toast.success('Uploaded Attachments Successfully')
          }
        } catch (error) {
          if (isAxiosError<ApiResponse>(error)) {
            if (error?.response?.data?.message) {
              toast.error(error?.response?.data?.message)
            } else {
              toast.error(error.message)
            }
          }
        }
      }
    }
  }
  return (
    <>
      <Button
        className={className}
        icon={icon}
        rounded={rounded}
        text
        size={size}
        label={label}
        onClick={handleButtonClick}
      />
      <input
        multiple={multiple}
        accept='image/*,video/*'
        ref={fileInputRef}
        type='file'
        className='hidden'
        onChange={handleFileChange}
      />
    </>
  )
}
export default FileUpload
