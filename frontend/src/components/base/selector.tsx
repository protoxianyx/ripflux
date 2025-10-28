import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

const Selector = () => {
  return (
    <Select>
      <SelectTrigger className="w-[180px]">
        <SelectValue placeholder="Quality" />
      </SelectTrigger>
      <SelectContent>
        <SelectItem value="light">Select 1</SelectItem>
        <SelectItem value="dark">Seleect 2</SelectItem>
        <SelectItem value="system">Select 3</SelectItem>
      </SelectContent>
    </Select>
  );
};

export default Selector;
