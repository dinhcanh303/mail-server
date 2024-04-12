import { ARRAY_VIDEO, END_POINT_MINIO } from '@/constants/const'
import { Attachment } from '@/models/Attachment'
import { Galleria } from 'primereact/galleria'
import { Image } from 'primereact/image'

interface GalleryProps {
  attachments?: Attachment[]
}
const Gallery: React.FC<GalleryProps> = ({ attachments }) => {
  const itemTemplate = (attachment: Attachment) => {
    if (ARRAY_VIDEO.includes(attachment?.extension)) {
      return (
        <video
          // onMouseOver={handleMouseOver}
          // onMouseOut={handleMouseOut}
          className='w-full h-full object-cover'
          controls
        >
          <source src={END_POINT_MINIO + attachment?.url} type={attachment?.mimeType} />
        </video>
      )
    }
    return <Image src={END_POINT_MINIO + attachment.url} alt={attachment.fileName} />
  }
  return (
    <div className='w-full h-full flex justify-center items-center p-2 rounded-lg bg-gray-200'>
      <Galleria value={attachments} circular showThumbnails={false} showItemNavigators item={itemTemplate} />
    </div>
  )
}

export default Gallery
