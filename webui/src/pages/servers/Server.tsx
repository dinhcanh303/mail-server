import { createServer, duplicateServer, getServers } from '@/apis/server.api'
import { Server as ModelServer } from '@/models/Server'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { BreadCrumb } from 'primereact/breadcrumb'
import { Button } from 'primereact/button'
import { Column } from 'primereact/column'
import { DataTable } from 'primereact/datatable'
import { Dialog } from 'primereact/dialog'
import { Dropdown, DropdownChangeEvent } from 'primereact/dropdown'
import { InputNumber, InputNumberValueChangeEvent } from 'primereact/inputnumber'
import { InputSwitch, InputSwitchChangeEvent } from 'primereact/inputswitch'
import { InputText } from 'primereact/inputtext'
import { Tag } from 'primereact/tag'
import { Toast } from 'primereact/toast'
import { Toolbar } from 'primereact/toolbar'
import { classNames } from 'primereact/utils'
import React, { useEffect, useRef, useState } from 'react'
import { FloatLabel } from 'primereact/floatlabel'

// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface ServerProps {}

// eslint-disable-next-line no-empty-pattern
const Server: React.FC<ServerProps> = ({}) => {
  const items = [{ label: 'Servers' }]
  const home = { icon: 'pi pi-home', url: '/' }

  const emptyServer: ModelServer = {
    name: '',
    host: 'smtp.yoursite.com',
    port: '465',
    authProtocol: 'plain',
    username: 'username',
    password: 'password',
    fromName: 'email',
    fromAddress: 'noreply@server.yoursite.com',
    tlsType: 'TLS',
    tlsSkipVerify: false,
    maxConnections: 10,
    idleTimeout: 15,
    retries: 5,
    waitTimeout: 10
  }
  const [servers, setServers] = useState(null)
  const [toEmail, setToEmail] = useState<string>('')
  const [serverDialog, setServerDialog] = useState(false)
  const [dupServerDialog, setDupServerDialog] = useState(false)
  const [testConnection, setTestConnection] = useState(false)
  const [deleteServerDialog, setDeleteServerDialog] = useState(false)
  const [deleteServersDialog, setDeleteServersDialog] = useState(false)
  const [server, setServer] = useState<ModelServer>(emptyServer)
  const [selectedServers, setSelectedServers] = useState<ModelServer | null>(null)
  const [submitted, setSubmitted] = useState(false)
  const [globalFilter, setGlobalFilter] = useState(null)
  const toast = useRef(null)
  const dt = useRef(null)
  const tlsTypes = [
    { name: 'Off', value: 'OFF' },
    { name: 'STARTTLS', value: 'STARTTLS' },
    { name: 'SSL/TLS', value: 'TLS' }
  ]
  const authProtocols = [
    { name: 'Plain', value: 'plain' },
    // { name: 'Login', value: 'login' },
    { name: 'CRAM-MD5', value: 'cram-md5' }
  ]
  const queryClient = useQueryClient()
  const serverRes = useQuery({
    queryKey: ['servers'],
    queryFn: () => {
      return getServers(10, 0)
    }
  })
  const tmp = serverRes.data?.data.servers
  useEffect(() => {
    setServers(tmp)
  }, [tmp])
  const openNew = () => {
    setServer(emptyServer)
    setSubmitted(false)
    setServerDialog(true)
  }

  const hideDialog = () => {
    setSubmitted(false)
    setTestConnection(false)
    setServerDialog(false)
  }

  const hideDeleteServerDialog = () => {
    setDeleteServerDialog(false)
  }

  const hideDeleteServersDialog = () => {
    setDeleteServersDialog(false)
  }
  const hideDuplicateServerDialog = () => {
    setDupServerDialog(false)
  }

  const handleCreateServer = async () => {
    setSubmitted(true)
    if (server.name.trim()) {
      try {
        console.log(server)
        const res = await createServer({ server: server })
        if (res?.status == 200) {
          toast?.current?.show({ severity: 'success', summary: 'Success', detail: 'Server Created Successfully' })
          setTestConnection(false)
          setServerDialog(false)
          setServer(emptyServer)
        } else toast?.current?.show({ severity: 'warning', summary: 'Warning', detail: 'Server Created Failed' })
      } catch (error) {
        console.log(error)
      }
    }
  }
  const handleDuplicateServer = async () => {
    try {
      const res = await duplicateServer({
        server: server
      })
      if (res?.status == 200) {
        toast?.current?.show({ severity: 'success', summary: 'Success', detail: 'Client Duplicated Successfully' })
        queryClient.invalidateQueries({ queryKey: [`servers`], exact: true })
        setDupServerDialog(false)
        setServer(emptyServer)
      } else toast?.current?.show({ severity: 'warning', summary: 'Warning', detail: 'Client Duplicated Failed' })
    } catch (error) {
      toast?.current?.show({ severity: 'error', summary: 'Failed', detail: 'Client Updated Failed', life: 3000 })
    }
  }

  const editServer = (server: ModelServer) => {
    setServer({ ...server })
    setTestConnection(false)
    setServerDialog(true)
  }
  const cloneClient = (server: ModelServer) => {
    server.name = 'Copy of ' + server.name
    setServer({ ...server })
    setDupServerDialog(true)
  }

  const confirmDeleteServer = (rowData: ModelServer) => {
    // setProduct(product)
    setDeleteServerDialog(true)
  }
  const onInputChange = (e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>, name: string) => {
    const val = (e.target && e.target.value) || ''
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const _server: any = { ...server }
    _server[name] = val
    setServer(_server)
  }

  const onInputNumberChange = (e: InputNumberValueChangeEvent, name: string) => {
    const val = e.value || 0
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const _server: any = { ...server }
    _server[name] = val
    setServer(_server)
  }
  const onDropdownChange = (e: DropdownChangeEvent, name: string) => {
    const val = e.value || 0
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const _server: any = { ...server }
    _server[name] = val
    setServer(_server)
  }
  const onInputSwitchChange = (e: InputSwitchChangeEvent, name: string) => {
    const val = e.value || 0
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const _server: any = { ...server }
    _server[name] = val
    setServer(_server)
  }

  const leftToolbarTemplate = () => {
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
      <h4 className='m-0'>Servers</h4>
      <span className='p-input-icon-left flex'>
        <InputText type='search' onInput={(e) => setGlobalFilter(e.target?.value)} placeholder='Search...' />
      </span>
    </div>
  )
  const serverDialogFooter = (
    <>
      <Button label='Cancel' icon='pi pi-times' outlined onClick={hideDialog} />
      <Button label='Save' icon='pi pi-check' onClick={handleCreateServer} />
    </>
  )
  const deleteServerDialogFooter = (
    <>
      <Button label='No' icon='pi pi-times' outlined onClick={hideDeleteServerDialog} />
      {/* <Button label='Yes' icon='pi pi-check' severity='danger' onClick={deleteProduct} /> */}
    </>
  )
  const deleteServersDialogFooter = (
    <>
      <Button label='No' icon='pi pi-times' outlined onClick={hideDeleteServersDialog} />
      {/* <Button label='Yes' icon='pi pi-check' severity='danger' onClick={deleteSelectedProducts} /> */}
    </>
  )
  const duplicateServerDialogFooter = (
    <>
      <Button label='Clone' icon='pi pi-check' onClick={handleDuplicateServer} />
    </>
  )
  const statusBodyTemplate = (rowData: ModelServer) => {
    switch (rowData.isDefault) {
      case true:
        return <Tag value='default' severity='info'></Tag>
      default:
        return ''
    }
  }
  const tlsTypeBodyTemplate = (rowData: ModelServer) => {
    return <Tag value={rowData.tlsType} severity={rowData.tlsType == 'TLS' ? 'success' : 'waring'}></Tag>
  }
  const actionBodyTemplate = (rowData: ModelServer) => {
    return (
      <div>
        <Button icon='pi pi-user-edit' size='small' rounded outlined text onClick={() => editServer(rowData)} />
        <Button icon='pi pi-clone' size='small' rounded outlined text onClick={() => cloneClient(rowData)} />
        <Button
          icon='pi pi-trash'
          size='small'
          rounded
          outlined
          text
          severity={rowData.isDefault ? 'secondary' : 'danger'}
          hidden
          onClick={rowData.isDefault ? undefined : () => confirmDeleteServer(rowData)}
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
            value={servers}
            selection={selectedServers}
            onSelectionChange={(e) => setSelectedServers(e.value as ModelServer)}
            dataKey='id'
            paginator
            rows={10}
            size='small'
            rowsPerPageOptions={[5, 10, 25]}
            paginatorTemplate='FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown'
            currentPageReportTemplate='Showing {first} to {last} of {totalRecords} servers'
            globalFilter={globalFilter}
            header={header}
          >
            <Column selectionMode='multiple' exportable={false}></Column>
            <Column field='id' header='ID' sortable style={{ minWidth: '5rem' }}></Column>
            <Column header='Action' body={actionBodyTemplate} exportable={false} style={{ minWidth: '12rem' }}></Column>
            <Column field='name' header='Name' sortable style={{ minWidth: '5rem' }}></Column>
            <Column
              field='isDefault'
              header='Default'
              sortable
              body={statusBodyTemplate}
              style={{ minWidth: '5rem' }}
            ></Column>
            <Column field='host' header='Host' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='port' header='Port' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='authProtocol' header='Auth Protocol' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='username' header='Username' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='password' header='Password' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='fromName' header='From Name' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='fromAddress' header='From Address' sortable style={{ minWidth: '10rem' }}></Column>
            <Column
              field='tlsType'
              header='TLS Type'
              body={tlsTypeBodyTemplate}
              sortable
              style={{ minWidth: '10rem' }}
            ></Column>
            <Column field='tlsSkipVerify' header='Skip Verification' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='maxConnections' header='Max Connections' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='idleTimeout' header='IDLE Timeout' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='retries' header='Retries' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='waitTimeout' header='Wait Timeout' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='createdAt' header='Created At' sortable style={{ minWidth: '12rem' }}></Column>
            <Column field='updatedAt' header='Updated At' sortable style={{ minWidth: '12rem' }}></Column>
          </DataTable>
        </div>

        <Dialog
          visible={serverDialog}
          style={{ width: '64rem' }}
          breakpoints={{ '960px': '75vw', '641px': '90vw' }}
          header='Server Details'
          modal
          className='p-fluid'
          footer={serverDialogFooter}
          onHide={hideDialog}
        >
          <div className='grid md:grid-cols-2 gap-x-5 gap-y-2'>
            <div className='field'>
              <label htmlFor='name' className='font-bold'>
                Name
              </label>
              <InputText
                id='name'
                value={server.name}
                onChange={(e) => onInputChange(e, 'name')}
                required
                autoFocus
                className={classNames({ 'p-invalid': submitted && !server.name })}
              />
              {submitted && !server.name && <small className='p-error'>Name is required.</small>}
            </div>
            <div className='field'>
              <label htmlFor='host' className='font-bold'>
                Host
              </label>
              <InputText id='host' value={server.host} onChange={(e) => onInputChange(e, 'host')} required />
            </div>
            <div className='field'>
              <label htmlFor='port' className='font-bold'>
                Port
              </label>
              <InputText id='port' value={server.port} onChange={(e) => onInputChange(e, 'port')} required />
            </div>
            <div className='field'>
              <label htmlFor='authProtocol' className='font-bold'>
                Auth Protocol
              </label>
              <Dropdown
                id='authProtocol'
                value={server.authProtocol}
                onChange={(e) => onDropdownChange(e, 'authProtocol')}
                options={authProtocols}
                optionLabel='name'
                optionValue='value'
                placeholder='Select a Auth Protocol'
                className='w-full md:w-14rem'
              />
            </div>
            <div className='field'>
              <label htmlFor='username' className='font-bold'>
                Username
              </label>
              <InputText
                id='username'
                value={server.username}
                onChange={(e) => onInputChange(e, 'username')}
                required
              />
            </div>
            <div className='field'>
              <label htmlFor='password' className='font-bold'>
                Password
              </label>
              <InputText
                id='password'
                value={server.password}
                onChange={(e) => onInputChange(e, 'password')}
                required
              />
            </div>
            <div className='field'>
              <label htmlFor='fromName' className='font-bold'>
                From Name
              </label>
              <InputText
                id='fromName'
                value={server.fromName}
                onChange={(e) => onInputChange(e, 'fromName')}
                required
              />
            </div>
            <div className='field'>
              <label htmlFor='fromAddress' className='font-bold'>
                From Address
              </label>
              <InputText
                id='fromAddress'
                value={server.fromAddress}
                onChange={(e) => onInputChange(e, 'fromAddress')}
                required
              />
            </div>
            <div className='flex gap-x-5'>
              <div className='field'>
                <label htmlFor='maxConnections' className='font-bold'>
                  Max Connections
                </label>
                <InputNumber
                  id='maxConnections'
                  value={server.maxConnections}
                  onValueChange={(e) => onInputNumberChange(e, 'maxConnections')}
                  mode='decimal'
                  min={1}
                  max={100}
                  showButtons
                  buttonLayout='horizontal'
                  incrementButtonIcon='pi pi-plus'
                  decrementButtonIcon='pi pi-minus'
                />
              </div>
              <div className='field'>
                <label htmlFor='idleTimeout' className='font-bold'>
                  IDLE Timeout
                </label>
                <InputNumber
                  id='idleTimeout'
                  value={server.idleTimeout}
                  onValueChange={(e) => onInputNumberChange(e, 'idleTimeout')}
                  mode='decimal'
                  min={5}
                  max={1000}
                  showButtons
                  buttonLayout='horizontal'
                  incrementButtonIcon='pi pi-plus'
                  decrementButtonIcon='pi pi-minus'
                />
              </div>
            </div>
            <div className='flex gap-x-5'>
              <div className='field'>
                <label htmlFor='retries' className='font-bold'>
                  Retries
                </label>
                <InputNumber
                  id='retries'
                  value={server.retries}
                  onValueChange={(e) => onInputNumberChange(e, 'retries')}
                  mode='decimal'
                  min={1}
                  max={100}
                  showButtons
                  buttonLayout='horizontal'
                  incrementButtonIcon='pi pi-plus'
                  decrementButtonIcon='pi pi-minus'
                />
              </div>
              <div className='field'>
                <label htmlFor='waitTimeout' className='font-bold'>
                  Wait Timeout
                </label>
                <InputNumber
                  id='waitTimeout'
                  value={server.waitTimeout}
                  onValueChange={(e) => onInputNumberChange(e, 'waitTimeout')}
                  mode='decimal'
                  min={1}
                  max={100}
                  showButtons
                  buttonLayout='horizontal'
                  incrementButtonIcon='pi pi-plus'
                  decrementButtonIcon='pi pi-minus'
                />
              </div>
            </div>
            <div className='field'>
              <label htmlFor='tlsSkipVerify' className='font-bold'>
                Skip Verification
              </label>
              <div className='flex justify-center items-center'>
                <InputSwitch
                  id='tlsSkipVerify'
                  checked={server.tlsSkipVerify}
                  onChange={(e) => onInputSwitchChange(e, 'tlsSkipVerify')}
                ></InputSwitch>
              </div>
            </div>
            <div className='field'>
              <label htmlFor='tlsType' className='font-bold'>
                TLS Type
              </label>
              <Dropdown
                id='tlsType'
                value={server.tlsType}
                onChange={(e) => onDropdownChange(e, 'tlsType')}
                options={tlsTypes}
                optionLabel='name'
                optionValue='value'
                placeholder='Select a Tls Type'
                className='w-full md:w-14rem'
              />
            </div>
          </div>
          <div
            className={`mt-3 border-t border-b cursor-pointer ${testConnection ? '' : 'flex justify-end'}`}
            onClick={() => setTestConnection(true)}
          >
            <span className={testConnection ? `hidden` : `text-blue-600 mt-3 mb-3`}>
              <i className='pi pi-bolt'></i>
              Test connection
            </span>
            <div className={testConnection ? 'flex justify-between mt-6 mb-6' : 'hidden'}>
              <div>
                <span className='font-bold'>Default from email</span>
                <p>{server.fromName + '<' + server.fromAddress + '>'}</p>
              </div>
              <FloatLabel>
                <InputText
                  id='toEmail'
                  value={toEmail}
                  onChange={(e: React.ChangeEvent<HTMLInputElement>) => setToEmail(e.target.value)}
                />
                <label htmlFor='toEmail'>To Email</label>
              </FloatLabel>
              <div>
                <Button label='Send Mail' icon='pi pi-check' onClick={handleDuplicateServer} />
              </div>
            </div>
          </div>
        </Dialog>

        <Dialog
          visible={deleteServerDialog}
          style={{ width: '32rem' }}
          breakpoints={{ '960px': '75vw', '641px': '90vw' }}
          header='Confirm'
          modal
          footer={deleteServerDialogFooter}
          onHide={hideDeleteServerDialog}
        >
          <div className='confirmation-content'>
            <i className='pi pi-exclamation-triangle mr-3' style={{ fontSize: '2rem' }} />
            {server && (
              <span>
                Are you sure you want to delete <b>{server.name}</b>?
              </span>
            )}
          </div>
        </Dialog>

        <Dialog
          visible={deleteServersDialog}
          style={{ width: '32rem' }}
          breakpoints={{ '960px': '75vw', '641px': '90vw' }}
          header='Confirm'
          modal
          footer={deleteServersDialogFooter}
          onHide={hideDeleteServersDialog}
        >
          <div className='confirmation-content'>
            <i className='pi pi-exclamation-triangle mr-3' style={{ fontSize: '2rem' }} />
            {server && <span>Are you sure you want to delete the selected products?</span>}
          </div>
        </Dialog>
        <Dialog
          visible={dupServerDialog}
          style={{ width: '32rem' }}
          breakpoints={{ '960px': '75vw', '641px': '90vw' }}
          header='Clone Client'
          modal
          footer={duplicateServerDialogFooter}
          onHide={hideDuplicateServerDialog}
        >
          <div className='flex justify-between items-center'>
            <label htmlFor='name' className='font-bold'>
              Name
            </label>
            <InputText
              id='name'
              value={server.name}
              onChange={(e) => onInputChange(e, 'name')}
              required
              autoFocus
              className={classNames({ 'p-invalid': submitted && !server.name })}
            />
            {submitted && !server.name && <small className='p-error'>Name is required.</small>}
          </div>
        </Dialog>
      </div>
    </div>
  )
}
export default Server
