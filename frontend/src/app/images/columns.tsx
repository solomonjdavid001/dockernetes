"use client";

import { ColumnDef } from "@tanstack/react-table";

export type Image = {
  name: string;
  tag: string;
  created_at: number;
  size: number;
  id: string;
};

export const columns: ColumnDef<Image>[] = [
  {
    accessorKey: "name",
    header: () => <div className="text-left">Name</div>,
    cell: ({ row }) => {
      const formatted: string = row.getValue("name");
      return <div className="text-left font-medium">{formatted}</div>;
    },
  },
  {
    accessorKey: "tag",
    header: "Tag",
  },
  {
    accessorKey: "created_at",
    header: "Created At",
  },
  {
    accessorKey: "size",
    header: "Size",
  },
  {
    accessorKey: "id",
    header: "ID",
  },
];
