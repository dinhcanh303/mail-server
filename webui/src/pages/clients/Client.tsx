import { createClient, deleteClient, duplicateClient, getClients, updateClient } from '@/apis/client.api'
import { getServers } from '@/apis/server.api'
import { getTemplatesActive } from '@/apis/template.api'
import { Client as ModelClient } from '@/models/Client'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { BreadCrumb } from 'primereact/breadcrumb'
import { Button } from 'primereact/button'
import { Column } from 'primereact/column'
import { DataTable } from 'primereact/datatable'
import { Dialog } from 'primereact/dialog'
import { Dropdown, DropdownChangeEvent } from 'primereact/dropdown'
import { InputText } from 'primereact/inputtext'
import { Tag } from 'primereact/tag'
import { Toast } from 'primereact/toast'
import { Toolbar } from 'primereact/toolbar'
import { classNames } from 'primereact/utils'
import React, { useEffect, useRef, useState } from 'react'

// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface ClientProps {}

// eslint-disable-next-line no-empty-pattern
const Client: React.FC<ClientProps> = ({}) => {
  const items = [{ label: 'Clients' }]
  const home = { icon: 'pi pi-home', url: '/' }

  const emptyClient: ModelClient = {
    name: '',
    serverId: '',
    templateId: ''
  }
  const [clients, setClients] = useState<ModelClient[] | null>(null)
  const [clientDialog, setClientDialog] = useState(false)
  const [dupClientDialog, setDupClientDialog] = useState(false)
  const [deleteClientDialog, setDeleteClientDialog] = useState(false)
  const [deleteClientsDialog, setDeleteClientsDialog] = useState(false)
  const [client, setClient] = useState<ModelClient>(emptyClient)
  const [selectedClients, setSelectedClients] = useState<ModelClient | null>(null)
  const [submitted, setSubmitted] = useState(false)
  const [globalFilter, setGlobalFilter] = useState(null)
  const toast = useRef(null)
  const dt = useRef(null)
  const queryClient = useQueryClient()
  const clientRes = useQuery({
    queryKey: ['clients'],
    queryFn: () => {
      return getClients(10, 0)
    }
  })
  const templatesActive = useQuery({
    queryKey: ['templates_active'],
    queryFn: () => {
      return getTemplatesActive()
    }
  }).data?.data.templates
  const serversActive = useQuery({
    queryKey: ['servers_active'],
    queryFn: () => {
      return getServers(1000, 0)
    }
  }).data?.data.servers
  const tmp = clientRes.data?.data.clients
  useEffect(() => {
    setClients(tmp)
  }, [tmp])
  const openNew = () => {
    setClient(emptyClient)
    setSubmitted(false)
    setClientDialog(true)
  }

  const hideDialog = () => {
    setSubmitted(false)
    setClientDialog(false)
  }
  const hideDeleteClientDialog = () => {
    setDeleteClientDialog(false)
  }

  const hideDeleteClientsDialog = () => {
    setDeleteClientsDialog(false)
  }
  const hideDuplicateClientDialog = () => {
    setDupClientDialog(false)
  }
  const handleDeleteClient = async () => {
    if (client.id) {
      try {
        const res = await deleteClient(client.id)
        if (res.status == 200) {
          queryClient.invalidateQueries({ queryKey: ['Clients'], exact: true })
          setDeleteClientDialog(false)
          setClient(emptyClient)
          toast?.current?.show({ severity: 'success', summary: 'Success', detail: 'Client Deleted', life: 3000 })
        }
      } catch (error) {
        toast?.current?.show({ severity: 'error', summary: 'Failed', detail: 'Client Deleted Failed', life: 3000 })
      }
    }
  }
  const handleCreateOrUpdateClient = async () => {
    setSubmitted(true)
    if (client.id) handleUpdateClient()
    else handleCreateClient()
  }
  const handleCreateClient = async () => {
    if (client.name.trim()) {
      try {
        const res = await createClient({
          client: client
        })
        if (res?.status == 200) {
          toast?.current?.show({ severity: 'success', summary: 'Success', detail: 'Client Created Successfully' })
          queryClient.invalidateQueries({ queryKey: [`Clients`], exact: true })
          setClientDialog(false)
          setClient(emptyClient)
        } else toast?.current?.show({ severity: 'warning', summary: 'Warning', detail: 'Client Created Failed' })
      } catch (error) {
        toast?.current?.show({ severity: 'error', summary: 'Failed', detail: 'Client Created Failed', life: 3000 })
      }
    }
  }

  const handleUpdateClient = async () => {
    try {
      const res = await updateClient({
        client: client
      })
      if (res?.status == 200) {
        toast?.current?.show({ severity: 'success', summary: 'Success', detail: 'Client Updated Successfully' })
        queryClient.invalidateQueries({ queryKey: [`clients`], exact: true })
        setClientDialog(false)
        setClient(emptyClient)
      } else toast?.current?.show({ severity: 'warning', summary: 'Warning', detail: 'Client Updated Failed' })
    } catch (error) {
      toast?.current?.show({ severity: 'error', summary: 'Failed', detail: 'Client Updated Failed', life: 3000 })
    }
  }
  const handleDuplicateClient = async () => {
    try {
      const res = await duplicateClient({
        client: client
      })
      if (res?.status == 200) {
        toast?.current?.show({ severity: 'success', summary: 'Success', detail: 'Client Duplicated Successfully' })
        queryClient.invalidateQueries({ queryKey: [`clients`], exact: true })
        setDupClientDialog(false)
        setClient(emptyClient)
      } else toast?.current?.show({ severity: 'warning', summary: 'Warning', detail: 'Client Duplicated Failed' })
    } catch (error) {
      toast?.current?.show({ severity: 'error', summary: 'Failed', detail: 'Client Updated Failed', life: 3000 })
    }
  }

  const editClient = (client: ModelClient) => {
    setClient({ ...client })
    setClientDialog(true)
  }
  const cloneClient = (client: ModelClient) => {
    client.name = 'Copy of ' + client.name
    setClient({ ...client })
    setDupClientDialog(true)
  }
  const confirmDeleteClient = (rowData: ModelClient) => {
    setClient(rowData)
    setDeleteClientDialog(true)
  }
  const onInputChange = (e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>, name: string) => {
    const val = (e.target && e.target.value) || ''
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const _client: any = { ...client }
    _client[name] = val
    setClient(_client)
  }
  const onDropdownChange = (e: DropdownChangeEvent, name: string) => {
    const val = e.value || 0
    console.log(e)
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const _client: any = { ...client }
    _client[name] = val
    setClient(_client)
  }

  const leftToolbarClient = () => {
    return (
      <div className='flex flex-wrap gap-2'>
        <Button label='New' size='small' icon='pi pi-plus' severity='success' onClick={openNew} />
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
  const header = (
    <div className='flex flex-wrap gap-2 align-items-center justify-between'>
      <h4 className='m-0'>Clients</h4>
      <span className='p-input-icon-left flex'>
        <InputText type='search' onInput={(e) => setGlobalFilter(e.target?.value)} placeholder='Search...' />
      </span>
    </div>
  )
  const clientDialogFooter = (
    <>
      <Button label='Cancel' icon='pi pi-times' outlined onClick={hideDialog} />
      <Button label='Save' icon='pi pi-check' onClick={handleCreateOrUpdateClient} />
    </>
  )
  const deleteClientDialogFooter = (
    <>
      <Button label='No' icon='pi pi-times' outlined onClick={hideDeleteClientDialog} />
      <Button label='Yes' icon='pi pi-check' severity='danger' onClick={handleDeleteClient} />
    </>
  )
  const deleteClientsDialogFooter = (
    <>
      <Button label='No' icon='pi pi-times' outlined onClick={hideDeleteClientsDialog} />
      {/* <Button label='Yes' icon='pi pi-check' severity='danger' onClick={deleteSelectedProducts} /> */}
    </>
  )
  const duplicateClientDialogFooter = (
    <>
      <Button label='Clone' icon='pi pi-check' onClick={handleDuplicateClient} />
    </>
  )
  const isDefaultBodyClient = (rowData: ModelClient) => {
    switch (rowData.isDefault) {
      case true:
        return <Tag value='default' severity='info'></Tag>
      default:
        return ''
    }
  }
  const actionBodyClient = (rowData: ModelClient) => {
    return (
      <div>
        <Button icon='pi pi-user-edit' size='small' rounded outlined text onClick={() => editClient(rowData)} />
        <Button icon='pi pi-clone' size='small' rounded outlined text onClick={() => cloneClient(rowData)} />
        <Button
          icon='pi pi-trash'
          size='small'
          rounded
          outlined
          text
          severity={rowData.isDefault ? 'secondary' : 'danger'}
          hidden
          onClick={rowData.isDefault ? undefined : () => confirmDeleteClient(rowData)}
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
          <Toolbar className='mb-4' left={leftToolbarClient}></Toolbar>
          <DataTable
            ref={dt}
            value={clients}
            selection={selectedClients}
            onSelectionChange={(e) => setSelectedClients(e.value as ModelClient)}
            dataKey='id'
            paginator
            rows={10}
            size='small'
            rowsPerPageOptions={[5, 10, 25]}
            paginatorClient='FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown'
            currentPageReportClient='Showing {first} to {last} of {totalRecords} Clients'
            globalFilter={globalFilter}
            header={header}
          >
            <Column selectionMode='multiple' exportable={false}></Column>
            <Column field='id' header='ID' sortable style={{ minWidth: '5rem' }}></Column>
            <Column header='Action' body={actionBodyClient} exportable={false} style={{ minWidth: '12rem' }}></Column>
            <Column field='name' header='Name' sortable style={{ minWidth: '5rem' }}></Column>
            <Column
              field='isDefault'
              header='Default'
              sortable
              body={isDefaultBodyClient}
              style={{ minWidth: '5rem' }}
            ></Column>
            <Column field='apiKey' header='Api Key' sortable style={{ minWidth: '20rem' }}></Column>
            <Column field='templateId' header='Template ID' sortable style={{ minWidth: '20rem' }}></Column>
            <Column field='serverId' header='Server ID' sortable style={{ minWidth: '20rem' }}></Column>
            <Column field='createdAt' header='Created At' sortable style={{ minWidth: '12rem' }}></Column>
            <Column field='updatedAt' header='Updated At' sortable style={{ minWidth: '12rem' }}></Column>
          </DataTable>
        </div>
        <Dialog
          visible={clientDialog}
          style={{ width: '64rem' }}
          breakpoints={{ '960px': '75vw', '641px': '90vw' }}
          header='Client Details'
          modal
          className='p-fluid'
          footer={clientDialogFooter}
          onHide={hideDialog}
        >
          <div className='grid gap-y-2'>
            <div className='field'>
              <label htmlFor='name' className='font-bold'>
                Name
              </label>
              <InputText
                id='name'
                value={client.name}
                onChange={(e) => onInputChange(e, 'name')}
                required
                autoFocus
                className={classNames({ 'p-invalid': submitted && !Client.name })}
              />
              {submitted && !Client.name && <small className='p-error'>Name is required.</small>}
            </div>
            <div className='field'>
              <label htmlFor='serverId' className='font-bold'>
                Server
              </label>
              <Dropdown
                id='serverId'
                value={client.serverId}
                onChange={(e) => onDropdownChange(e, 'serverId')}
                options={serversActive}
                optionLabel='name'
                optionValue='id'
                placeholder='Select a server'
                className='w-full md:w-14rem'
              />
            </div>
            <div className='field'>
              <label htmlFor='templateId' className='font-bold'>
                Template
              </label>
              <Dropdown
                id='templateId'
                value={client.templateId}
                onChange={(e) => onDropdownChange(e, 'templateId')}
                options={templatesActive}
                optionLabel='name'
                optionValue='id'
                placeholder='Select a template'
                className='w-full md:w-14rem'
              />
            </div>
          </div>
        </Dialog>
        <Dialog
          visible={deleteClientDialog}
          style={{ width: '32rem' }}
          breakpoints={{ '960px': '75vw', '641px': '90vw' }}
          header='Confirm'
          modal
          footer={deleteClientDialogFooter}
          onHide={hideDeleteClientDialog}
        >
          <div className='confirmation-content'>
            <i className='pi pi-exclamation-triangle mr-3' style={{ fontSize: '2rem' }} />
            {Client && (
              <span>
                Are you sure you want to delete <b>{Client.name}</b>?
              </span>
            )}
          </div>
        </Dialog>
        <Dialog
          visible={deleteClientsDialog}
          style={{ width: '32rem' }}
          breakpoints={{ '960px': '75vw', '641px': '90vw' }}
          header='Confirm'
          modal
          footer={deleteClientsDialogFooter}
          onHide={hideDeleteClientsDialog}
        >
          <div className='confirmation-content'>
            <i className='pi pi-exclamation-triangle mr-3' style={{ fontSize: '2rem' }} />
            {client && <span>Are you sure you want to delete the selected products?</span>}
          </div>
        </Dialog>
        <Dialog
          visible={dupClientDialog}
          style={{ width: '32rem' }}
          breakpoints={{ '960px': '75vw', '641px': '90vw' }}
          header='Clone Client'
          modal
          footer={duplicateClientDialogFooter}
          onHide={hideDuplicateClientDialog}
        >
          <div className='flex justify-between items-center'>
            <label htmlFor='name' className='font-bold'>
              Name
            </label>
            <InputText
              id='name'
              value={client.name}
              onChange={(e) => onInputChange(e, 'name')}
              required
              autoFocus
              className={classNames({ 'p-invalid': submitted && !client.name })}
            />
            {submitted && !client.name && <small className='p-error'>Name is required.</small>}
          </div>
        </Dialog>
      </div>
    </div>
  )
}
export default Client
