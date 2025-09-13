import {
  Card,
  CardHeader,
  CardTitle,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";

export default function NotFoundPage() {
  return (
    <div className="flex min-h-screen items-center justify-center bg-background px-4">
      <Card className="max-w-md w-full text-center">
        <CardHeader>
          <CardTitle className="text-4xl font-bold text-destructive">
            404
          </CardTitle>
        </CardHeader>

        <CardContent className="space-y-2">
          <h2 className="text-2xl font-semibold">Page not found!</h2>
          <p className="text-muted-foreground">
            The page you are looking for does not exist.
          </p>
        </CardContent>

        <CardFooter className="flex justify-center">
          <Button asChild>
            <a href="/">Go to Home</a>
          </Button>
        </CardFooter>
      </Card>
    </div>
  );
}
