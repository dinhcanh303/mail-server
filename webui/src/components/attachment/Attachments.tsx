import { Attachment } from '@/models/Attachment'
import { ARRAY_IMAGE_VIDEO, ARRAY_VIDEO, END_POINT_MINIO } from '@/constants/const'
import { Image } from 'primereact/image'
import { useRef } from 'react'
import { MdOutlinePlayCircle } from 'react-icons/md'
interface AttachmentsProps {
  attachments?: Attachment[]
  setOpenPostModal?: any
}
const Attachments: React.FC<AttachmentsProps> = ({ attachments, setOpenPostModal }) => {
  if (attachments === undefined) return
  // eslint-disable-next-line react-hooks/rules-of-hooks
  const videoRef = useRef<HTMLVideoElement>(null)
  // // Function to handle mouseover event
  // const handleMouseOver = () => {
  //     if (videoRef.current) {
  //     videoRef.current.play()
  //     }
  // }
  // // Function to handle mouseout event
  // const handleMouseOut = () => {
  //     if (videoRef.current) {
  //     videoRef.current.pause()
  //     videoRef.current.currentTime = 0;
  //     }
  // }
  // const handleOpenGallery = (attachmentId: string) => {
  //     const a = document.getElementById(attachmentId)
  //     if(a) a.click()
  // }
  const attachmentIsNotImageAndVideo = attachments.filter((attachment) => {
    return !ARRAY_IMAGE_VIDEO.includes(attachment.extension)
  })
  let attachmentIsNotImageAndVideoRender
  let attachmentsRender
  if (attachmentIsNotImageAndVideo.length > 0) {
    attachmentIsNotImageAndVideoRender = attachmentIsNotImageAndVideo.map((attachment, _) => {
      return (
        <a key={attachment?.id} href={END_POINT_MINIO + attachment?.url}>
          {attachment?.fileName}
        </a>
      )
    })
  }
  if (attachments.length > 0) {
    let className = 'grid grid-cols-2'
    if (attachments.length === 1) className = 'flex justify-center'
    attachmentsRender = (
      <div className={className}>
        {attachments.slice(0, 3).map((attachment, _) => {
          if (ARRAY_VIDEO.includes(attachment?.extension)) {
            return (
              <div
                key={attachment?.id}
                className='relative hover:text-gray-200 text-gray-600 cursor-pointer'
                onClick={setOpenPostModal}
              >
                <div className='absolute top-1/2 left-1/2 -mt-6 -ml-6'>
                  <MdOutlinePlayCircle className='w-12 h-12' />
                </div>
                <video
                  ref={videoRef}
                  // onMouseOver={handleMouseOver}
                  // onMouseOut={handleMouseOut}
                  className='w-auto h-full object-cover'
                  src={END_POINT_MINIO + attachment?.url}
                />
              </div>
            )
          }
          return (
            <Image
              className='cursor-pointer'
              onClick={setOpenPostModal}
              key={attachment?.id}
              src={END_POINT_MINIO + attachment?.url}
              alt='Image'
              width='auto'
            />
          )
        })}
        {attachments.length > 3 && (
          <div className='relative bg-gray-100 opacity-50 w-auto h-auto cursor-pointer'>
            {/* <FcAddImage className='absolute top-1/2 left-1/2 -mt-10 -ml-10 w-20 h-20' /> */}
            <div className='absolute top-1/2 left-1/2 -mt-4 -ml-1 text-black font-semibold text-lg'>
              +{attachments.length - 3}
            </div>
            <Image
              className='flex justify-center'
              onClick={setOpenPostModal}
              src={END_POINT_MINIO + attachments[3]?.url}
              alt='Image'
              width='auto'
            />
          </div>
        )}
      </div>
    )
  }
  // const handleRenderByExtension = (attachment : Attachment) => {
  //     const url = END_POINT_MINIO + attachment?.url
  //     let urlThumbnail = END_POINT_MINIO + attachment?.urlThumbnail
  //     const name  = attachment?.fileName
  //     const extension = attachment?.extension
  //     let mimeType = attachment?.mimeType
  //     let dataIframe = false
  //     let dataSource
  //     switch (extension) {
  //         case 'MP4':
  //         case 'MOV':
  //             mimeType = 'video/mp4'
  //             break;
  //         default:
  //             break;
  //     }
  //     if(['video/mp4','video/webm'].includes(mimeType)){
  //         dataSource = '{"source": [{"src":"' + url +'","type":"'+mimeType+'"}], "attributes": {"preload": false, "controls": true}}'
  //         urlThumbnail = "src/assets/gallery/video.png"
  //     }
  //     return {name,urlThumbnail,url,dataIframe,dataSource}
  // }
  // const items = attachments.map((attachment,_) => {
  //     const {name,urlThumbnail,url,dataSource} = handleRenderByExtension(attachment)
  //     return (
  //         <a key={attachment?.id} id={attachment?.id} href={url}
  //         data-lg-size="1280-720"
  //         data-video={dataSource}
  //         data-poster={urlThumbnail}
  //         >
  //             <img alt={name} src={urlThumbnail} />
  //         </a>
  //     )
  // })
  return (
    <>
      {attachmentIsNotImageAndVideoRender}
      {attachmentsRender}
      {/* <div className='hidden'>
                <LightGallery
                autoplayFirstVideo={false}
                    pager={false}
                    speed={500}
                    plugins={[
                        lgZoom,
                        lgThumbnail,
                        lgRotate,
                        lgFullscreen,
                        lgAutoplay,
                        lgVideo,
                    ]}
                    mobileSettings={{
                        controls: false,
                        showCloseIcon: false,
                        download: false,
                        rotate: false
                    }}
                >
                    {items}
                </LightGallery>
            </div> */}
    </>
  )
}
export default Attachments
