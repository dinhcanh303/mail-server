import { getServers } from '@/apis/server.api'
import { Server as ModelServer } from '@/models/Server'
import { useQuery } from '@tanstack/react-query'
import { BreadCrumb } from 'primereact/breadcrumb'
import { Button } from 'primereact/button'
import { Column } from 'primereact/column'
import { DataTable } from 'primereact/datatable'
import { Dialog } from 'primereact/dialog'
import { Dropdown, DropdownChangeEvent } from 'primereact/dropdown'
import { InputNumber, InputNumberValueChangeEvent } from 'primereact/inputnumber'
import { InputSwitch, InputSwitchChangeEvent } from 'primereact/inputswitch'
import { InputText } from 'primereact/inputtext'
import { InputTextarea } from 'primereact/inputtextarea'
import { RadioButton } from 'primereact/radiobutton'
import { Tag } from 'primereact/tag'
import { Toast } from 'primereact/toast'
import { Toolbar } from 'primereact/toolbar'
import { classNames } from 'primereact/utils'
import React, { useEffect, useRef, useState } from 'react'

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
    username: 'username',
    password: 'password',
    tlsType: 'TLS',
    tlsSkipVerify: false,
    maxConnections: 10,
    idleTimeout: 15,
    retries: 5,
    waitTimeout: 10
  }
  const [servers, setServers] = useState(null)
  const [serverDialog, setServerDialog] = useState(false)
  const [deleteProductDialog, setDeleteProductDialog] = useState(false)
  const [deleteProductsDialog, setDeleteProductsDialog] = useState(false)
  const [server, setServer] = useState<ModelServer>(emptyServer)
  const [selectedProducts, setSelectedProducts] = useState(null)
  const [submitted, setSubmitted] = useState(false)
  const [globalFilter, setGlobalFilter] = useState(null)
  const toast = useRef(null)
  const dt = useRef(null)
  const tlsTypes = [
    { name: 'Off', value: 'OFF' },
    { name: 'STARTTLS', value: 'STARTTLS' },
    { name: 'SSL/TLS', value: 'TLS' }
  ]
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
    // setProduct(emptyProduct)
    setSubmitted(false)
    setServerDialog(true)
  }

  const hideDialog = () => {
    setSubmitted(false)
    setServerDialog(false)
  }

  const hideDeleteProductDialog = () => {
    setDeleteProductDialog(false)
  }

  const hideDeleteProductsDialog = () => {
    setDeleteProductsDialog(false)
  }

  // const saveProduct = () => {
  //   setSubmitted(true)

  //   if (product.name.trim()) {
  //     const _products = [...products]
  //     const _product = { ...product }

  //     if (product.id) {
  //       const index = findIndexById(product.id)

  //       _products[index] = _product
  //       toast.current.show({ severity: 'success', summary: 'Successful', detail: 'Product Updated', life: 3000 })
  //     } else {
  //       _product.id = createId()
  //       _product.image = 'product-placeholder.svg'
  //       _products.push(_product)
  //       toast.current.show({ severity: 'success', summary: 'Successful', detail: 'Product Created', life: 3000 })
  //     }

  //     setProducts(_products)
  //     setProductDialog(false)
  //     setProduct(emptyProduct)
  //   }
  // }

  const editServer = (server: ModelServer) => {
    setServer({ ...server })
    setServerDialog(true)
  }

  // const confirmDeleteProduct = (product) => {
  //   setProduct(product)
  //   setDeleteProductDialog(true)
  // }

  // const deleteProduct = () => {
  //   const _products = products.filter((val) => val.id !== product.id)

  //   setProducts(_products)
  //   setDeleteProductDialog(false)
  //   setProduct(emptyProduct)
  //   toast.current.show({ severity: 'success', summary: 'Successful', detail: 'Product Deleted', life: 3000 })
  // }

  // const findIndexById = (id) => {
  //   let index = -1

  //   for (let i = 0; i < products.length; i++) {
  //     if (products[i].id === id) {
  //       index = i
  //       break
  //     }
  //   }

  //   return index
  // }

  // const createId = () => {
  //   let id = ''
  //   const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'

  //   for (let i = 0; i < 5; i++) {
  //     id += chars.charAt(Math.floor(Math.random() * chars.length))
  //   }

  //   return id
  // }

  // const confirmDeleteSelected = () => {
  //   setDeleteProductsDialog(true)
  // }

  // const deleteSelectedProducts = () => {
  //   const _products = products.filter((val) => !selectedProducts.includes(val))

  //   setProducts(_products)
  //   setDeleteProductsDialog(false)
  //   setSelectedProducts(null)
  //   toast.current.show({ severity: 'success', summary: 'Successful', detail: 'Products Deleted', life: 3000 })
  // }

  // const onCategoryChange = (e) => {
  //   const _product = { ...product }

  //   _product['category'] = e.value
  //   setProduct(_product)
  // }

  const onInputChange = (e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>, name: string) => {
    const val = (e.target && e.target.value) || ''
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const _server: any = { ...server }
    _server[name] = val
    setServer(_server)
  }

  const onInputNumberChange = (e: InputNumberValueChangeEvent, name: string) => {
    console.log(name)
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
        {/* <i className='pi pi-search ml-2' /> */}
      </span>
    </div>
  )
  const serverDialogFooter = (
    <React.Fragment>
      <Button label='Cancel' icon='pi pi-times' outlined onClick={hideDialog} />
      {/* <Button label='Save' icon='pi pi-check' onClick={saveProduct} /> */}
    </React.Fragment>
  )
  const deleteProductDialogFooter = (
    <React.Fragment>
      <Button label='No' icon='pi pi-times' outlined onClick={hideDeleteProductDialog} />
      {/* <Button label='Yes' icon='pi pi-check' severity='danger' onClick={deleteProduct} /> */}
    </React.Fragment>
  )
  const deleteProductsDialogFooter = (
    <React.Fragment>
      <Button label='No' icon='pi pi-times' outlined onClick={hideDeleteProductsDialog} />
      {/* <Button label='Yes' icon='pi pi-check' severity='danger' onClick={deleteSelectedProducts} /> */}
    </React.Fragment>
  )
  const statusBodyTemplate = (rowData: ModelServer) => {
    console.log(rowData)
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
        <Button
          icon='pi pi-trash'
          size='small'
          rounded
          outlined
          text
          severity='danger'
          // onClick={() => confirmDeleteServer(rowData)}
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
            selection={selectedProducts}
            // onSelectionChange={(e) => setSelectedProducts(e.value)}
            dataKey='id'
            paginator
            rows={10}
            size='small'
            rowsPerPageOptions={[5, 10, 25]}
            paginatorTemplate='FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown'
            currentPageReportTemplate='Showing {first} to {last} of {totalRecords} products'
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
            <Column field='username' header='Username' sortable style={{ minWidth: '10rem' }}></Column>
            <Column field='password' header='Password' sortable style={{ minWidth: '10rem' }}></Column>
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
              <InputSwitch
                id='tlsSkipVerify'
                checked={server.tlsSkipVerify}
                onChange={(e) => onInputSwitchChange(e, 'tlsSkipVerify')}
              ></InputSwitch>
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
        </Dialog>

        <Dialog
          visible={deleteProductDialog}
          style={{ width: '32rem' }}
          breakpoints={{ '960px': '75vw', '641px': '90vw' }}
          header='Confirm'
          modal
          footer={deleteProductDialogFooter}
          onHide={hideDeleteProductDialog}
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
          visible={deleteProductsDialog}
          style={{ width: '32rem' }}
          breakpoints={{ '960px': '75vw', '641px': '90vw' }}
          header='Confirm'
          modal
          footer={deleteProductsDialogFooter}
          onHide={hideDeleteProductsDialog}
        >
          <div className='confirmation-content'>
            <i className='pi pi-exclamation-triangle mr-3' style={{ fontSize: '2rem' }} />
            {server && <span>Are you sure you want to delete the selected products?</span>}
          </div>
        </Dialog>
      </div>
    </div>
  )
}
export default Server
