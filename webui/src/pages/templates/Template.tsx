import { createTemplate, deleteTemplate, getTemplates, updateTemplate } from '@/apis/template.api'
import { Template as ModelTemplate } from '@/models/Template'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { BreadCrumb } from 'primereact/breadcrumb'
import { Button } from 'primereact/button'
import { Column } from 'primereact/column'
import { DataTable } from 'primereact/datatable'
import { Dialog } from 'primereact/dialog'
import { Dropdown, DropdownChangeEvent } from 'primereact/dropdown'
import { InputText } from 'primereact/inputtext'
import { InputTextarea } from 'primereact/inputtextarea'
import { Tag } from 'primereact/tag'
import { Toast } from 'primereact/toast'
import { Toolbar } from 'primereact/toolbar'
import { classNames } from 'primereact/utils'
import React, { useEffect, useRef, useState } from 'react'

// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface TemplateProps {}

// eslint-disable-next-line no-empty-pattern
const Template: React.FC<TemplateProps> = ({}) => {
  const items = [{ label: 'Templates' }]
  const home = { icon: 'pi pi-home', url: '/' }

  const emptyTemplate: ModelTemplate = {
    name: '',
    html: ''
  }
  const [templates, setTemplates] = useState<ModelTemplate[] | undefined>(undefined)
  const [templateDialog, setTemplateDialog] = useState(false)
  const [templatePreviewDialog, setTemplatePreviewDialog] = useState(false)
  const [deleteTemplateDialog, setDeleteTemplateDialog] = useState(false)
  const [deleteTemplatesDialog, setDeleteTemplatesDialog] = useState(false)
  const [template, setTemplate] = useState<ModelTemplate>(emptyTemplate)
  const [selectedTemplates, setSelectedTemplates] = useState<ModelTemplate[]>([])
  const [submitted, setSubmitted] = useState(false)
  const [globalFilter, setGlobalFilter] = useState<string>('')
  const toast = useRef(null)
  const dt = useRef(null)
  const statues = [
    { name: 'Active', value: 'active' },
    { name: 'Inactive', value: 'inactive' }
  ]
  const queryClient = useQueryClient()
  const templateRes = useQuery({
    queryKey: ['templates'],
    queryFn: () => {
      return getTemplates(10, 0)
    }
  })
  const tmp = templateRes.data?.data.templates
  useEffect(() => {
    setTemplates(tmp)
  }, [tmp])
  const openNew = () => {
    setTemplate(emptyTemplate)
    setSubmitted(false)
    setTemplateDialog(true)
  }

  const hideDialog = () => {
    setSubmitted(false)
    setTemplateDialog(false)
  }
  const hidePreviewDialog = () => {
    setTemplatePreviewDialog(false)
  }
  const hideDeleteTemplateDialog = () => {
    setDeleteTemplateDialog(false)
  }

  const hideDeleteTemplatesDialog = () => {
    setDeleteTemplatesDialog(false)
  }
  const handleDeleteTemplate = async () => {
    if (template.id) {
      try {
        const res = await deleteTemplate(template.id)
        if (res.status == 200) {
          queryClient.invalidateQueries({ queryKey: ['templates'], exact: true })
          setDeleteTemplateDialog(false)
          setTemplate(emptyTemplate)
          toast?.current?.show({ severity: 'success', summary: 'Success', detail: 'Template Deleted', life: 3000 })
        }
      } catch (error) {
        toast?.current?.show({ severity: 'error', summary: 'Failed', detail: 'Template Deleted Failed', life: 3000 })
      }
    }
  }
  const handleCreateOrUpdateTemplate = async () => {
    setSubmitted(true)
    if (template.id) handleUpdateTemplate()
    else handleCreateTemplate()
  }
  const handleCreateTemplate = async () => {
    if (template.name.trim()) {
      try {
        const res = await createTemplate({
          template: template
        })
        if (res?.status == 200) {
          toast?.current?.show({ severity: 'success', summary: 'Success', detail: 'Template Created Successfully' })
          queryClient.invalidateQueries({ queryKey: [`templates`], exact: true })
          setTemplateDialog(false)
          setTemplate(emptyTemplate)
        } else toast?.current?.show({ severity: 'warning', summary: 'Warning', detail: 'Template Created Failed' })
      } catch (error) {
        toast?.current?.show({ severity: 'error', summary: 'Failed', detail: 'Template Created Failed', life: 3000 })
      }
    }
  }

  const handleUpdateTemplate = async () => {
    try {
      const res = await updateTemplate({
        template: template
      })
      if (res?.status == 200) {
        toast?.current?.show({ severity: 'success', summary: 'Success', detail: 'Template Updated Successfully' })
        queryClient.invalidateQueries({ queryKey: [`templates`], exact: true })
        setTemplateDialog(false)
        setTemplate(emptyTemplate)
      } else toast?.current?.show({ severity: 'warning', summary: 'Warning', detail: 'Template Updated Failed' })
    } catch (error) {
      toast?.current?.show({ severity: 'error', summary: 'Failed', detail: 'Template Updated Failed', life: 3000 })
    }
  }

  const editTemplate = (template: ModelTemplate) => {
    setTemplate({ ...template })
    setTemplateDialog(true)
  }
  const previewTemplate = (template: ModelTemplate) => {
    setTemplate({ ...template })
    setTemplatePreviewDialog(true)
  }
  const confirmDeleteTemplate = (rowData: ModelTemplate) => {
    setTemplate(rowData)
    setDeleteTemplateDialog(true)
  }
  const onInputChange = (e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>, name: string) => {
    const val = (e.target && e.target.value) || ''
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const _template: any = { ...template }
    _template[name] = val
    setTemplate(_template)
  }
  const onDropdownChange = (e: DropdownChangeEvent, name: string) => {
    const val = e.value || 0
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const _template: any = { ...template }
    _template[name] = val
    setTemplate(_template)
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
      <h4 className='m-0'>Templates</h4>
      <span className='p-input-icon-left flex'>
        <InputText type='search' onInput={(e) => setGlobalFilter(e.currentTarget?.value)} placeholder='Search...' />
      </span>
    </div>
  )
  const templateDialogFooter = (
    <>
      <Button label='Cancel' icon='pi pi-times' outlined onClick={hideDialog} />
      <Button label='Save' icon='pi pi-check' onClick={handleCreateOrUpdateTemplate} />
    </>
  )
  const deleteTemplateDialogFooter = (
    <>
      <Button label='No' icon='pi pi-times' outlined onClick={hideDeleteTemplateDialog} />
      <Button label='Yes' icon='pi pi-check' severity='danger' onClick={handleDeleteTemplate} />
    </>
  )
  const deleteTemplatesDialogFooter = (
    <>
      <Button label='No' icon='pi pi-times' outlined onClick={hideDeleteTemplatesDialog} />
      {/* <Button label='Yes' icon='pi pi-check' severity='danger' onClick={deleteSelectedProducts} /> */}
    </>
  )
  const isDefaultBodyTemplate = (rowData: ModelTemplate) => {
    switch (rowData.isDefault) {
      case true:
        return <Tag value='default' severity='info'></Tag>
      default:
        return ''
    }
  }
  const statusBodyTemplate = (rowData: ModelTemplate) => {
    const type = rowData.status == 'active' ? 'success' : 'waring'
    return <Tag value={rowData.status} severity={type as 'success' | 'warning'}></Tag>
  }
  const actionBodyTemplate = (rowData: ModelTemplate) => {
    return (
      <div>
        <Button icon='pi pi-eye' size='small' rounded outlined text onClick={() => previewTemplate(rowData)} />
        <Button icon='pi pi-user-edit' size='small' rounded outlined text onClick={() => editTemplate(rowData)} />
        <Button
          icon='pi pi-trash'
          size='small'
          rounded
          outlined
          text
          severity={rowData.isDefault ? 'secondary' : 'danger'}
          hidden
          onClick={rowData.isDefault ? undefined : () => confirmDeleteTemplate(rowData)}
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
            value={templates}
            selection={selectedTemplates}
            onSelectionChange={(e) => {
              if (Array.isArray(e.value)) {
                setSelectedTemplates(e.value)
              }
            }}
            dataKey='id'
            paginator
            rows={10}
            size='small'
            rowsPerPageOptions={[5, 10, 25]}
            paginatorTemplate='FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown'
            currentPageReportTemplate='Showing {first} to {last} of {totalRecords} templates'
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
              body={isDefaultBodyTemplate}
              style={{ minWidth: '5rem' }}
            ></Column>
            <Column field='html' header='HTML' sortable style={{ minWidth: '20rem' }}></Column>
            <Column
              field='status'
              header='Status'
              body={statusBodyTemplate}
              sortable
              style={{ minWidth: '10rem' }}
            ></Column>
            <Column field='createdAt' header='Created At' sortable style={{ minWidth: '12rem' }}></Column>
            <Column field='updatedAt' header='Updated At' sortable style={{ minWidth: '12rem' }}></Column>
          </DataTable>
        </div>
        <Dialog
          visible={templateDialog}
          style={{ width: '64rem' }}
          breakpoints={{ '960px': '75vw', '641px': '90vw' }}
          header='Template Details'
          modal
          className='p-fluid'
          footer={templateDialogFooter}
          onHide={hideDialog}
        >
          <div className='grid gap-y-2'>
            <div className='field'>
              <label htmlFor='name' className='font-bold'>
                Name
              </label>
              <InputText
                id='name'
                value={template.name}
                onChange={(e) => onInputChange(e, 'name')}
                required
                autoFocus
                className={classNames({ 'p-invalid': submitted && !template.name })}
              />
              {submitted && !template.name && <small className='p-error'>Name is required.</small>}
            </div>
            <div className='field'>
              <label htmlFor='status' className='font-bold'>
                Status
              </label>
              <Dropdown
                id='status'
                value={template.status}
                onChange={(e) => onDropdownChange(e, 'status')}
                options={statues}
                optionLabel='name'
                optionValue='value'
                placeholder='Select a status'
                className='w-full md:w-14rem'
              />
            </div>
            <div className='field'>
              <div className='flex justify-between items-center'>
                <label htmlFor='html' className='font-bold'>
                  HTML
                </label>
                <div>
                  <Button
                    icon='pi pi-eye mr-1'
                    size='small'
                    rounded
                    outlined
                    onClick={() => previewTemplate(template)}
                    title='preview'
                  >
                    Preview
                  </Button>
                </div>
              </div>

              <InputTextarea
                id='html'
                value={template.html}
                onChange={(e) => onInputChange(e, 'html')}
                required
                rows={10}
              />
            </div>
          </div>
        </Dialog>
        <Dialog
          visible={templatePreviewDialog}
          style={{ width: '50rem' }}
          breakpoints={{ '960px': '75vw', '641px': '90vw' }}
          header='Preview'
          modal
          // footer={deleteTemplateDialogFooter}
          onHide={hidePreviewDialog}
        >
          <div className='border rounded-lg'>
            {template && <div dangerouslySetInnerHTML={{ __html: template.html }}></div>}
          </div>
        </Dialog>

        <Dialog
          visible={deleteTemplateDialog}
          style={{ width: '32rem' }}
          breakpoints={{ '960px': '75vw', '641px': '90vw' }}
          header='Confirm'
          modal
          footer={deleteTemplateDialogFooter}
          onHide={hideDeleteTemplateDialog}
        >
          <div className='confirmation-content'>
            <i className='pi pi-exclamation-triangle mr-3' style={{ fontSize: '2rem' }} />
            {template && (
              <span>
                Are you sure you want to delete <b>{template.name}</b>?
              </span>
            )}
          </div>
        </Dialog>
        <Dialog
          visible={deleteTemplatesDialog}
          style={{ width: '32rem' }}
          breakpoints={{ '960px': '75vw', '641px': '90vw' }}
          header='Confirm'
          modal
          footer={deleteTemplatesDialogFooter}
          onHide={hideDeleteTemplatesDialog}
        >
          <div className='confirmation-content'>
            <i className='pi pi-exclamation-triangle mr-3' style={{ fontSize: '2rem' }} />
            {template && <span>Are you sure you want to delete the selected products?</span>}
          </div>
        </Dialog>
      </div>
    </div>
  )
}
export default Template
