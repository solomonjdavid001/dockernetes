"use client";

import { Area, AreaChart, CartesianGrid, XAxis } from "recharts";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

// Initial chart data
const initialData = [{ time: "00:00", value: 0 }];

type ChartComponentProps = {
  metric: "cpu" | "memory" | "disk";
  title: string;
  description: string;
  color: string;
};

export function ChartComponent({
  title,
  description,
  color,
}: ChartComponentProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>{title}</CardTitle>
        <CardDescription>{description}</CardDescription>
      </CardHeader>
      <CardContent>
        <ChartContainer config={{ value: { label: title, color } }}>
          <AreaChart
            data={initialData}
            margin={{
              top: 10,
              right: 20,
              bottom: 20,
              left: 20,
            }}
          >
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="time"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              tickFormatter={(value) => value.slice(0, 5)}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="line" />}
            />
            <Area
              dataKey="value"
              type="natural"
              fill={color}
              fillOpacity={0.4}
              stroke={color}
            />
          </AreaChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
