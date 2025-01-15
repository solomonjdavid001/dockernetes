
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
    header: "Name",
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
