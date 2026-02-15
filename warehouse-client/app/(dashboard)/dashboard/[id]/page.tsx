'use client'
import useGetAllBarang from '@/hooks/useGetAllBarang'
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
import { useParams, useRouter } from 'next/navigation'
import useGetBarang from '@/hooks/useGetBarang'
import useGetStokByBarangID from '@/hooks/useGetStokByBarangID'
import useGetHistoryByBarangID from '@/hooks/useGetHistoryByBarangID'

const Page = () => {
    const params = useParams()
    const barangID = params.id

    const { barang, error, loading } = useGetBarang(Number(barangID))
    const { stok } = useGetStokByBarangID(Number(barangID))

    const { historyStoksByBarangID } = useGetHistoryByBarangID(Number(barangID))

    return (
        <div className="w-full min-h-screen flex">
            <div className='w-300 mx-auto justify-center'>
                <div className='text-3xl font-bold mt-12'>{barang?.nama_barang}</div>
                <Table className='bg-white rounded-lg w-full mt-4'>
                    <TableHeader>
                        <TableRow>
                            <TableHead className="w-25">Nama Barang</TableHead>
                            <TableHead>Kode Barang</TableHead>
                            <TableHead>Unit</TableHead>
                            <TableHead>Stok Akhir</TableHead>
                            <TableHead className="text-right">Harga Beli</TableHead>
                            <TableHead className="text-right">Harga Jual</TableHead>
                            <TableHead className="text-right">Updated At</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        <TableRow className='cursor-pointer' key={barang?.id}>
                            <TableCell className="font-medium">{barang?.nama_barang}</TableCell>
                            <TableCell>{barang?.kode_barang}</TableCell>
                            <TableCell>{barang?.satuan}</TableCell>
                            <TableCell>{stok?.stok_akhir}</TableCell>
                            <TableCell className="text-right">Rp. {barang?.harga_beli}</TableCell>
                            <TableCell className="text-right">Rp. {barang?.harga_jual}</TableCell>
                            <TableCell className="text-right">{stok?.updated_at.split("T")[0]}</TableCell>
                        </TableRow>
                    </TableBody>
                </Table>


                {/* HISTORY */}
                <div className='text-3xl font-bold mt-12'>History Barang</div>
                <Table className='bg-white rounded-lg w-full mt-4'>
                    {/* <TableCaption>A list of your recent invoices.</TableCaption> */}
                    <TableHeader>
                        <TableRow>
                            <TableHead>Nama Barang</TableHead>
                            <TableHead>Kode Barang</TableHead>
                            <TableHead>Jumlah</TableHead>
                            <TableHead>Unit</TableHead>
                            <TableHead className="w-[100px]">Keterangan</TableHead>
                            <TableHead>Stok Sebelum</TableHead>
                            <TableHead>Stok Sesudah</TableHead>
                            <TableHead>Jenis Transaksi</TableHead>
                            <TableHead className="text-right">Nama Staff</TableHead>
                            <TableHead className="text-right">Updated At</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {historyStoksByBarangID && historyStoksByBarangID.map((historyStok) => (
                            <TableRow key={historyStok.id}>
                                <TableCell className="font-medium">{historyStok.barang.nama_barang}</TableCell>
                                <TableCell>{historyStok.barang.kode_barang}</TableCell>
                                <TableCell>{historyStok.jumlah}</TableCell>
                                <TableCell>{historyStok.barang.satuan}</TableCell>
                                <TableCell>{historyStok.jenis_transaksi}</TableCell>
                                <TableCell>{historyStok.stok_sebelum}</TableCell>
                                <TableCell>{historyStok.stok_sesudah}</TableCell>
                                <TableCell>{historyStok.keterangan}</TableCell>
                                <TableCell className="text-right">{historyStok.user.full_name}</TableCell>
                                <TableCell className="text-right space-x-4">{historyStok.created_at.split("T")[0] + " | " + historyStok.created_at.split("T")[1].slice(0, 8)}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </div>
        </div>
    )
}

export default Page