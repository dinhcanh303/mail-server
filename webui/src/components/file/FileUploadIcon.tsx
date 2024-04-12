import { FileUpload, ItemTemplateOptions } from 'primereact/fileupload'
import { Tooltip } from 'primereact/tooltip'

interface FileUploadIconProps {}

const FileUploadIcon: React.FC<FileUploadIconProps> = ({}) => {
  const chooseOptions = {
    icon: 'ml-1 pi pi-camera',
    iconOnly: true,
    className: 'rounded-full hover:bg-gray-300 p-1 text-black border-none bg-transparent'
  }

  return (
    <>
      <FileUpload
        className='target-emoji-file-upload-icon cursor-pointer'
        data-pr-tooltip='Attach a photo or video'
        data-pr-position='top'
        mode='advanced'
        name='demo[]'
        url='/api/upload'
        accept='image/*,video/*'
        maxFileSize={1000000}
        chooseOptions={chooseOptions}
      />
      <Tooltip className='text-xs' target='.target-emoji-file-upload-icon' />
    </>
  )
}

export default FileUploadIcon
