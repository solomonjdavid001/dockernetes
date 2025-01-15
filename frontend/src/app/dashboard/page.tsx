import { ChartComponent } from "@/components/chart";

export default function DashboardPage() {
  return (
    <div className="grid auto-rows-min gap-4 md:grid-cols-3">
      <ChartComponent
        metric="cpu"
        title="CPU Usage"
        description="Real-time CPU usage percentage"
        color="hsl(var(--chart-1))"
      />
      <ChartComponent
        metric="memory"
        title="Memory Usage"
        description="Real-time memory usage percentage"
        color="hsl(var(--chart-2))"
      />
      <ChartComponent
        metric="disk"
        title="Disk Usage"
        description="Real-time disk usage percentage"
        color="hsl(var(--chart-3))"
      />
    </div>
  );
}
