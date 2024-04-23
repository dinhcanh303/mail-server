import { getHistories } from '@/apis/history.api'
import { History as ModelHistory } from '@/models/History'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { BreadCrumb } from 'primereact/breadcrumb'
import { Button } from 'primereact/button'
import { Column } from 'primereact/column'
import { DataTable } from 'primereact/datatable'
import { Dialog } from 'primereact/dialog'
import { InputText } from 'primereact/inputtext'
import { Tag } from 'primereact/tag'
import { Toast } from 'primereact/toast'
import { Toolbar } from 'primereact/toolbar'
import React, { useEffect, useRef, useState } from 'react'
import { object } from 'yup'

// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface HistoryProps {}

// eslint-disable-next-line no-empty-pattern
const History: React.FC<HistoryProps> = ({}) => {
  const items = [{ label: 'Histories' }]
  const home = { icon: 'pi pi-home', url: '/' }

  const emptyHistory: ModelHistory = {
    apiKey: '',
    to: '',
    bcc: '',
    cc: '',
    subject: '',
    content: {},
    status: ''
  }
  const [histories, setHistories] = useState<ModelHistory[] | undefined>(undefined)
  const [toEmail, setToEmail] = useState<string>('')
  const [historyDialog, setHistoryDialog] = useState(false)
  const [testConnection, setTestConnection] = useState(false)
  const [deleteHistoryDialog, setDeleteHistoryDialog] = useState(false)
  const [deleteHistoriesDialog, setDeleteHistoriesDialog] = useState(false)
  const [history, setHistory] = useState<ModelHistory>(emptyHistory)
  const [selectedHistories, setSelectedHistories] = useState<ModelHistory[]>([])
  const [submitted, setSubmitted] = useState(false)
  const [globalFilter, setGlobalFilter] = useState<string>('')
  const toast = useRef(null)
  const dt = useRef(null)
  const queryClient = useQueryClient()
  const historiesRes = useQuery({
    queryKey: ['histories'],
    queryFn: () => {
      return getHistories(10, 0)
    }
  })
  const tmp = historiesRes.data?.data.histories
  useEffect(() => {
    setHistories(tmp)
  }, [tmp])
  const openNew = () => {
    // setHistory(emptyHistorie)
    setSubmitted(false)
    // setHistoriesDialog(true)
  }

  const hideDialog = () => {
    setSubmitted(false)
    setTestConnection(false)
    setHistoryDialog(false)
  }

  const hideDeleteHistoryDialog = () => {
    setDeleteHistoryDialog(false)
  }

  const hideDeleteHistoriesDialog = () => {
    setDeleteHistoriesDialog(false)
  }
  const retryHistory = (rowData: ModelHistory) => {
    console.log(rowData)
  }
  const confirmDeleteHistory = (rowData: ModelHistory) => {
    setDeleteHistoryDialog(true)
  }
  const header = (
    <div className='flex flex-wrap gap-2 align-items-center justify-between'>
      <h4 className='m-0'>Histories</h4>
      <span className='p-input-icon-left flex'>
        <InputText type='search' onInput={(e) => setGlobalFilter(e.currentTarget?.value)} placeholder='Search...' />
      </span>
    </div>
  )
  const historieDialogFooter = (
    <>
      <Button label='Cancel' icon='pi pi-times' outlined onClick={hideDialog} />
      {/* <Button label='Save' icon='pi pi-check' onClick={handleCreateHistorie} /> */}
    </>
  )
  const statusBodyTemplate = (rowData: ModelHistory) => {
    switch (rowData.status) {
      case 'failed':
        return <Tag value={rowData.status} severity='danger'></Tag>
      case 'success':
        return <Tag value={rowData.status} severity='success'></Tag>
      case 'pending':
      default:
        return <Tag value={rowData.status} severity='warning'></Tag>
    }
  }
  const contentBodyTemplate = (rowData: ModelHistory) => {
    return <div>{JSON.stringify(rowData.content)}</div>
  }
  const actionBodyTemplate = (rowData: ModelHistory) => {
    return (
      <div>
        <Button icon='pi pi-user-edit' size='small' rounded outlined text onClick={() => retryHistory(rowData)} />
        <Button
          icon='pi pi-trash'
          size='small'
          rounded
          outlined
          text
          severity='danger'
          hidden
          onClick={() => confirmDeleteHistory(rowData)}
        />
      </div>
    )
  }
  const deleteHistoryDialogFooter = (
    <>
      <Button label='No' icon='pi pi-times' outlined onClick={hideDeleteHistoryDialog} />
      {/* <Button label='Yes' icon='pi pi-check' severity='danger' onClick={deleteProduct} /> */}
    </>
  )
  const leftToolbarTemplate = () => {
    return (
      <div className='flex flex-wrap gap-2'>
        <Button
          label='Delete'
          size='small'
          icon='pi pi-trash'
          severity='danger'
          // onClick={confirmDeleteSelected}
          // disabled={!selectedProducts || !selectedProducts.length}
        />
      </div>
    )
  }
  return (
    <div className='p-2'>
      <BreadCrumb model={items} home={home} />
      <div>
        <Toast ref={toast} />
        <div className='card'>
          <Toolbar className='mb-4' left={leftToolbarTemplate}></Toolbar>
          <DataTable
            ref={dt}
            value={histories}
            selection={selectedHistories}
            onSelectionChange={(e) => {
              if (Array.isArray(e.value)) {
                setSelectedHistories(e.value)
              }
            }}
            dataKey='id'
            paginator
            rows={10}
            size='small'
            rowsPerPageOptions={[5, 10, 25]}
            paginatorTemplate='FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown'
            currentPageReportTemplate='Showing {first} to {last} of {totalRecords} histories'
            globalFilter={globalFilter}
            header={header}
          >
            <Column selectionMode='multiple' exportable={false}></Column>
            <Column field='id' header='ID' sortable style={{ minWidth: '5rem' }}></Column>
            <Column header='Action' body={actionBodyTemplate} exportable={false} style={{ minWidth: '12rem' }}></Column>
            <Column field='apiKey' header='Api Key' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='to' header='To' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='subject' header='Subject' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='cc' header='Cc' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='bcc' header='Bcc' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='content' header='Content' body={contentBodyTemplate} style={{ minWidth: '10rem' }}></Column>
            <Column
              field='status'
              header='Status'
              sortable
              body={statusBodyTemplate}
              style={{ minWidth: '5rem' }}
            ></Column>
            <Column field='createdAt' header='Created At' sortable style={{ minWidth: '12rem' }}></Column>
            <Column field='updatedAt' header='Updated At' sortable style={{ minWidth: '12rem' }}></Column>
          </DataTable>
        </div>
      </div>
      <Dialog
        visible={deleteHistoryDialog}
        style={{ width: '32rem' }}
        breakpoints={{ '960px': '75vw', '641px': '90vw' }}
        header='Confirm'
        modal
        footer={deleteHistoryDialogFooter}
        onHide={hideDeleteHistoryDialog}
      >
        <div className='confirmation-content'>
          <i className='pi pi-exclamation-triangle mr-3' style={{ fontSize: '2rem' }} />
          {history && (
            <span>
              Are you sure you want to delete <b>{history.id}</b>?
            </span>
          )}
        </div>
      </Dialog>
    </div>
  )
}
export default History
