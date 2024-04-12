import { upload } from '@/apis/upload.api'
import { ApiResponse } from '@/types/app'
import { isAxiosError } from 'axios'
import { FileUpload, FileUploadHandlerEvent, FileUploadRemoveEvent, FileUploadUploadEvent } from 'primereact/fileupload'
import toast from 'react-hot-toast'

interface FileUploadCustomProps {
  name?: string
  auto?: boolean
  multiple?: boolean
  accept?: string
  maxFileSize?: number
  onUploadComplete?: any
}

const FileUploadCustom: React.FC<FileUploadCustomProps> = ({
  name = 'files',
  auto = false,
  multiple = true,
  accept = 'image/*,video/*',
  maxFileSize = 1000000000,
  onUploadComplete
}) => {
  const emptyTemplate = () => {
    return (
      <div className='flex items-center align-items-center flex-col'>
        <div>
          <i className='text-green-500 bg-gray-100 pi pi-image mt-2 p-5 rounded-full border'></i>
        </div>
        <span className='my-5 font-light text-base italic'>Drag and Drop Image Here</span>
      </div>
    )
  }
  const uploadHandler = async (event: FileUploadHandlerEvent) => {
    const files: File[] = event.files
    const formData = new FormData()
    const name = event?.options?.props?.name ?? 'files'
    for (let i = 0; i < files.length; i++) {
      formData.append(name, files[i])
    }
    try {
      const res = await upload(formData)
      if (res?.status == 200) {
        event.options.props.onUpload
        const ids = res.data.map((item, _) => {
          return item.id
        })
        onUploadComplete(ids)
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
  // const onRemove = (event: FileUploadRemoveEvent) => {
  //     console.log(event)

  // }
  // const onUpload = (event: FileUploadUploadEvent) => {
  //     console.log(event)

  // }
  // const onClear = () => {
  //     console.log(123)
  // }
  const chooseOptions = {
    icon: 'pi pi-fw pi-images',
    iconOnly: true,
    className: 'custom-choose-btn w-10 h-10 p-button-rounded p-button-outlined'
  }
  const uploadOptions = {
    icon: 'pi pi-fw pi-cloud-upload',
    iconOnly: true,
    className: 'custom-upload-btn w-10 h-10 p-button-success p-button-rounded p-button-outlined'
  }
  const cancelOptions = {
    icon: 'pi pi-fw pi-times',
    iconOnly: true,
    className: 'custom-cancel-btn w-10 h-10 p-button-danger p-button-rounded p-button-outlined'
  }
  return (
    <FileUpload
      name={name}
      customUpload
      uploadHandler={uploadHandler}
      accept={accept}
      maxFileSize={maxFileSize}
      multiple={multiple}
      headerClassName='p-2'
      chooseOptions={chooseOptions}
      uploadOptions={uploadOptions}
      cancelOptions={cancelOptions}
      // onRemove={onRemove}
      // onClear={onClear}
      // onUpload={onUpload}
      auto={auto}
      emptyTemplate={emptyTemplate}
    />
  )
}
export default FileUploadCustom
