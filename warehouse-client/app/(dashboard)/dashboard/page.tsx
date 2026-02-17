'use client'
import useGetAllBarang, { MasterBarang } from '@/hooks/useGetAllBarang'
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"
import { Button } from '@/components/ui/button'
import { Edit, Trash } from 'lucide-react'
import {
    Pagination,
    PaginationContent,
    PaginationItem,
    PaginationLink,
    PaginationNext,
    PaginationPrevious,
} from "@/components/ui/pagination"
import {
    Dialog,
    DialogClose,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog"
import { Field, FieldGroup } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Textarea } from '@/components/ui/textarea'
import useCreateBarang from '@/hooks/useCreateBarang'
import { useState } from 'react'
import toast from 'react-hot-toast'
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"
import useDeleteBarang from '@/hooks/useDeleteBarang'
import {
    AlertDialog,
    AlertDialogAction,
    AlertDialogCancel,
    AlertDialogContent,
    AlertDialogDescription,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogTrigger,
} from "@/components/ui/alert-dialog"
import Link from 'next/link'
import { useRouter } from 'next/navigation'
import useUpdateBarang from '@/hooks/useUpdateBarang'
import { Barang } from '@/hooks/useGetAllHistoryStok'

const items = [
    { label: "Pilih satuan", value: null },
    { label: "unit", value: "unit" },
    { label: "pcs", value: "pcs" },
]


const Page = () => {
    const { barangs, meta, setPageNum, pageNum, limitNum, setBarangs, search, setSearch } = useGetAllBarang()
    const totalPage = Math.ceil((meta?.total || 0) / limitNum)

    const { createBarang, error, loading, setError } = useCreateBarang()
    const { updateBarang } = useUpdateBarang()
    const [barangID, setBarangID] = useState(0)
    const [namaBarang, setNamaBarang] = useState("")
    const [deskripsi, setDeskripsi] = useState("")
    const [satuan, setSatuan] = useState("")
    const [hargaBeli, setHargaBeli] = useState(0)
    const [hargaJual, setHargaJual] = useState(0)


    const { deleteBarang } = useDeleteBarang()

    const router = useRouter()
    const handleSubmit = async (e: React.SubmitEvent) => {
        e.preventDefault()

        if (!satuan || satuan == "") {
            setError("Satuan tidak boleh kosong")
            toast.error("Satuan tidak boleh kosong")
            return
        }

        const res = await createBarang(namaBarang, deskripsi, satuan, hargaBeli, hargaJual)
        if (res?.success) {
            toast.success("Successfully create barang")
        } else {
            toast.error(res?.message)
        }
        setBarangs(prev => [res?.data, ...prev])
        setNamaBarang("")
        setDeskripsi("")
        setSatuan("")
        setHargaBeli(0)
        setHargaJual(0)
    }

    const handleUpdateBarangSubmit = async (e: React.SubmitEvent) => {
        e.preventDefault()

        if (!satuan || satuan == "") {
            setError("Satuan tidak boleh kosong")
            toast.error("Satuan tidak boleh kosong")
            return
        }

        const res = await updateBarang(barangID, namaBarang, deskripsi, satuan, hargaBeli, hargaJual)
        if (res?.success) {
            toast.success("Successfully update barang")
        } else {
            toast.error(res?.message)
        }

        setBarangs(prev => prev.map(item => item.id == barangID ? {
            ...item,
            nama_barang: namaBarang,
            deskripsi,
            satuan,
            harga_beli: hargaBeli,
            harga_jual: hargaJual,
        } : item))

    }

    const handleEditClick = (barang: MasterBarang) => {
        setBarangID(barang.id)
        setNamaBarang(barang.nama_barang)
        setDeskripsi(barang.deskripsi)
        setSatuan(barang.satuan)
        setHargaBeli(barang.harga_beli)
        setHargaJual(barang.harga_jual)
    }

    return (
        <div className="w-full min-h-screen flex">
            <div className='w-300 mx-auto justify-center'>
                <div className='text-3xl font-bold mt-12'>DASHBOARD</div>
                <Dialog>
                    <div className='flex gap-4'>
                        <Input
                            placeholder='Search'
                            className='max-w-60 cursor-text bg-white mt-4 hover:bg-white/40 font-semibold py-2 px-3 rounded-lg text-sm shadow-sm shadow-gray-300'
                            value={search} onChange={e => setSearch(e.target.value)}
                        />
                        <DialogTrigger className='cursor-pointer bg-white mt-4 hover:bg-white/40 font-semibold py-2 px-3 rounded-lg text-sm shadow-sm shadow-gray-300'>
                            Create Barang
                        </DialogTrigger>

                    </div>
                    <DialogContent className="sm:max-w-sm">
                        <form onSubmit={handleSubmit} className=''>
                            <DialogHeader>
                                <DialogTitle>Create Barang</DialogTitle>
                                <DialogDescription>
                                    Masukkan detail barang di kolom berikut
                                </DialogDescription>
                            </DialogHeader>
                            <FieldGroup className='gap-2 mt-3'>
                                <Field>
                                    <Label htmlFor="nama_barang">Nama Barang</Label>
                                    <Input id="nama_barang" name="nama_barang" value={namaBarang} onChange={e => setNamaBarang(e.target.value)} required />
                                </Field>
                                <Field>
                                    <Label htmlFor="deskripsi">Deskripsi</Label>
                                    <Textarea id="deskripsi" name="deskripsi" value={deskripsi} onChange={e => setDeskripsi(e.target.value)} />
                                </Field>
                                <Field>
                                    <Label htmlFor="deskripsi">Satuan</Label>
                                    <Select value={satuan} onValueChange={(value) => setSatuan(value)} required>
                                        <SelectTrigger className="w-full">
                                            <SelectValue placeholder="Pilih Satuan" />
                                        </SelectTrigger>
                                        <SelectContent>
                                            <SelectGroup>
                                                <SelectLabel>Fruits</SelectLabel>
                                                {items.map((item) => (
                                                    <SelectItem key={item.value} value={item.value!}>
                                                        {item.label}
                                                    </SelectItem>
                                                ))}
                                            </SelectGroup>
                                        </SelectContent>
                                    </Select>
                                </Field>
                                <Field>
                                    <Label htmlFor="harga_beli">Harga Beli</Label>
                                    <Input id="harga_beli" name="harga_beli" type='number' min={0} value={hargaBeli} onChange={e => setHargaBeli(e.target.valueAsNumber)} required />
                                </Field>
                                <Field>
                                    <Label htmlFor="harga_jual">Harga Jual</Label>
                                    <Input id="harga_jual" name="harga_jual" type='number' min={0} value={hargaJual} onChange={e => setHargaJual(e.target.valueAsNumber)} required />
                                </Field>
                            </FieldGroup>
                            <DialogFooter className='mt-2'>
                                <DialogClose className='cursor-pointer bg-white hover:bg-white/40 font-semibold py-2 px-3 rounded-lg text-sm shadow-sm shadow-gray-300'>
                                    Cancel
                                </DialogClose>
                                <Button className='cursor-pointer' type="submit">Submit</Button>
                            </DialogFooter>
                        </form>
                    </DialogContent>
                </Dialog>
                <Table className='bg-white rounded-lg w-full mt-2'>
                    <TableHeader>
                        <TableRow>
                            <TableHead className="w-25">Nama Barang</TableHead>
                            <TableHead>Kode Barang</TableHead>
                            <TableHead>Unit</TableHead>
                            <TableHead className="text-right">Harga Beli</TableHead>
                            <TableHead className="text-right">Harga Jual</TableHead>
                            <TableHead className="text-right"></TableHead>
                            <TableHead className="text-left">Action</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {barangs && barangs.map((barang) => (
                            <TableRow onClick={() => router.replace(`dashboard/${barang.id}`)} className='cursor-pointer' key={barang.id}>
                                <TableCell className="font-medium">{barang.nama_barang}</TableCell>
                                <TableCell>{barang.kode_barang}</TableCell>
                                <TableCell>{barang.satuan}</TableCell>
                                <TableCell className="text-right">Rp. {barang.harga_beli}</TableCell>
                                <TableCell className="text-right">Rp. {barang.harga_jual}</TableCell>
                                <TableCell className="text-right">
                                    <Dialog>
                                        <DialogTrigger
                                            onClick={async (e) => {
                                                e.stopPropagation()
                                                handleEditClick(barang)
                                            }}
                                            className='cursor-pointer bg-gray-200 hover:bg-blue-500  text-blue-500 hover:text-white font-semibold w-8 h-8 rounded-sm text-sm shadow-sm shadow-gray-300'>
                                            <Edit className='w-5 mx-auto' />
                                        </DialogTrigger>

                                        <DialogContent className="sm:max-w-sm" onClick={e => e.stopPropagation()}>
                                            <form onSubmit={(e) => {
                                                handleUpdateBarangSubmit(e)
                                            }} className=''>
                                                <DialogHeader>
                                                    <DialogTitle>Update Barang</DialogTitle>
                                                    <DialogDescription>
                                                        Masukkan detail barang di kolom berikut
                                                    </DialogDescription>
                                                </DialogHeader>
                                                <FieldGroup className='gap-2 mt-3'>
                                                    <Field>
                                                        <Label htmlFor="nama_barang">Nama Barang</Label>
                                                        <Input id="nama_barang" name="nama_barang" value={namaBarang} onChange={e => setNamaBarang(e.target.value)} required />
                                                    </Field>
                                                    <Field>
                                                        <Label htmlFor="deskripsi">Deskripsi</Label>
                                                        <Textarea id="deskripsi" name="deskripsi" value={deskripsi} onChange={e => setDeskripsi(e.target.value)} />
                                                    </Field>
                                                    <Field>
                                                        <Label htmlFor="deskripsi">Satuan</Label>
                                                        <Select value={satuan} onValueChange={(value) => setSatuan(value)} required>
                                                            <SelectTrigger className="w-full">
                                                                <SelectValue placeholder="Pilih Satuan" />
                                                            </SelectTrigger>
                                                            <SelectContent>
                                                                <SelectGroup>
                                                                    <SelectLabel>Fruits</SelectLabel>
                                                                    {items.map((item) => (
                                                                        <SelectItem key={item.value} value={item.value!}>
                                                                            {item.label}
                                                                        </SelectItem>
                                                                    ))}
                                                                </SelectGroup>
                                                            </SelectContent>
                                                        </Select>
                                                    </Field>
                                                    <Field>
                                                        <Label htmlFor="harga_beli">Harga Beli</Label>
                                                        <Input id="harga_beli" name="harga_beli" type='number' min={0} value={hargaBeli} onChange={e => setHargaBeli(e.target.valueAsNumber)} required />
                                                    </Field>
                                                    <Field>
                                                        <Label htmlFor="harga_jual">Harga Jual</Label>
                                                        <Input id="harga_jual" name="harga_jual" type='number' min={0} value={hargaJual} onChange={e => setHargaJual(e.target.valueAsNumber)} required />
                                                    </Field>
                                                </FieldGroup>
                                                <DialogFooter className='mt-2'>
                                                    <DialogClose className='cursor-pointer bg-white hover:bg-white/40 font-semibold py-2 px-3 rounded-lg text-sm shadow-sm shadow-gray-300'>
                                                        Cancel
                                                    </DialogClose>
                                                    <Button className='cursor-pointer' type="submit">Submit</Button>
                                                </DialogFooter>
                                            </form>
                                        </DialogContent>
                                    </Dialog>
                                </TableCell>
                                <TableCell className='text-center'>
                                    <AlertDialog>
                                        <AlertDialogTrigger asChild>

                                            <Button className='bg-gray-200 rounded-sm hover:bg-red-500 cursor-pointer hover:text-white text-red-500' onClick={async (e) => {
                                                e.stopPropagation()
                                            }}>
                                                <Trash className='' />
                                            </Button>
                                        </AlertDialogTrigger >

                                        <AlertDialogContent onClick={async (e) => {
                                            e.stopPropagation()
                                        }}>
                                            <AlertDialogHeader>
                                                <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
                                                <AlertDialogDescription>
                                                    This action cannot be undone. This will delete the goods
                                                </AlertDialogDescription>
                                            </AlertDialogHeader>
                                            <AlertDialogFooter>
                                                <AlertDialogCancel className="cursor-pointer" onClick={async (e) => {
                                                    e.stopPropagation()
                                                }}>Cancel</AlertDialogCancel>
                                                <AlertDialogAction className="bg-red-500 cursor-pointer" onClick={async (e) => {
                                                    e.stopPropagation()
                                                    const res = await deleteBarang(barang.id)
                                                    setBarangs(prev => [...prev].filter(p => p.id != barang.id))
                                                    toast.success(res)
                                                }}>
                                                    Continue
                                                </AlertDialogAction>
                                            </AlertDialogFooter>
                                        </AlertDialogContent>
                                    </AlertDialog>
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>

                </Table>
                <Pagination className='mt-4'>
                    <PaginationContent>
                        <PaginationItem>
                            <PaginationPrevious className='cursor-pointer' onClick={() => pageNum > 1 && setPageNum(pageNum - 1)} />
                        </PaginationItem>
                        {
                            Array.from({ length: totalPage }).map((_, i) => (
                                <PaginationItem className='cursor-pointer' key={i}>
                                    <PaginationLink onClick={() => {
                                        setPageNum(i + 1)
                                    }} isActive={pageNum == i + 1}>
                                        {i + 1}
                                    </PaginationLink>
                                </PaginationItem>
                            ))
                        }
                        <PaginationItem>
                            <PaginationNext className='cursor-pointer' onClick={() => pageNum != totalPage && setPageNum(pageNum + 1)} />
                        </PaginationItem>
                    </PaginationContent>
                </Pagination>


            </div>
        </div>
    )
}

export default Page



