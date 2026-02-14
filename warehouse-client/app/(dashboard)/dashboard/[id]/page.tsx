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

const Page = () => {
    const params = useParams()
    const barangID = params.id

    const { barang, error, loading } = useGetBarang(Number(barangID))
    const { stok } = useGetStokByBarangID(Number(barangID))

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
            </div>
        </div>
    )
}

export default Page