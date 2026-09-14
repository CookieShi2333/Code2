<?php declare(strict_types=1);
namespace PHPUnit\Framework;
use function function_exists;
use ArrayAccess;
use Countable;
use PHPUnit\Framework\Constraint\ArrayHasKey;
use PHPUnit\Framework\Constraint\Callback;
use PHPUnit\Framework\Constraint\Constraint;
use PHPUnit\Framework\Constraint\Count;
use PHPUnit\Framework\Constraint\DirectoryExists;
use PHPUnit\Framework\Constraint\FileExists;
use PHPUnit\Framework\Constraint\GreaterThan;
use PHPUnit\Framework\Constraint\IsAnything;
use PHPUnit\Framework\Constraint\IsEmpty;
use PHPUnit\Framework\Constraint\IsEqual;
use PHPUnit\Framework\Constraint\IsEqualCanonicalizing;
use PHPUnit\Framework\Constraint\IsEqualIgnoringCase;
use PHPUnit\Framework\Constraint\IsEqualWithDelta;
use PHPUnit\Framework\Constraint\IsFalse;
use PHPUnit\Framework\Constraint\IsFinite;
use PHPUnit\Framework\Constraint\IsIdentical;
use PHPUnit\Framework\Constraint\IsInfinite;
use PHPUnit\Framework\Constraint\IsInstanceOf;
use PHPUnit\Framework\Constraint\IsJson;
use PHPUnit\Framework\Constraint\IsList;
use PHPUnit\Framework\Constraint\IsNan;
use PHPUnit\Framework\Constraint\IsNull;
use PHPUnit\Framework\Constraint\IsReadable;
use PHPUnit\Framework\Constraint\IsTrue;
use PHPUnit\Framework\Constraint\IsType;
use PHPUnit\Framework\Constraint\IsWritable;
use PHPUnit\Framework\Constraint\LessThan;
use PHPUnit\Framework\Constraint\LogicalAnd;
use PHPUnit\Framework\Constraint\LogicalNot;
use PHPUnit\Framework\Constraint\LogicalOr;
use PHPUnit\Framework\Constraint\LogicalXor;
use PHPUnit\Framework\Constraint\ObjectEquals;
use PHPUnit\Framework\Constraint\RegularExpression;
use PHPUnit\Framework\Constraint\StringContains;
use PHPUnit\Framework\Constraint\StringEndsWith;
use PHPUnit\Framework\Constraint\StringEqualsStringIgnoringLineEndings;
use PHPUnit\Framework\Constraint\StringEqualsStringIgnoringWhitespace;
use PHPUnit\Framework\Constraint\StringMatchesFormatDescription;
use PHPUnit\Framework\Constraint\StringStartsWith;
use PHPUnit\Framework\Constraint\TraversableContainsEqual;
use PHPUnit\Framework\Constraint\TraversableContainsIdentical;
use PHPUnit\Framework\Constraint\TraversableContainsOnly;
use PHPUnit\Framework\MockObject\Rule\AnyInvokedCount as AnyInvokedCountMatcher;
use PHPUnit\Framework\MockObject\Rule\InvokedAtLeastCount as InvokedAtLeastCountMatcher;
use PHPUnit\Framework\MockObject\Rule\InvokedAtLeastOnce as InvokedAtLeastOnceMatcher;
use PHPUnit\Framework\MockObject\Rule\InvokedAtMostCount as InvokedAtMostCountMatcher;
use PHPUnit\Framework\MockObject\Rule\InvokedCount as InvokedCountMatcher;
use PHPUnit\Framework\MockObject\Stub\Exception as ExceptionStub;
use PHPUnit\Util\Xml\XmlException;
use Throwable;
if (!function_exists('PHPUnit\Framework\assertArrayIsEqualToArrayOnlyConsideringListOfKeys')) {
    function assertArrayIsEqualToArrayOnlyConsideringListOfKeys(array $expected, array $actual, array $keysToBeConsidered, string $message = ''): void
    {
        Assert::assertArrayIsEqualToArrayOnlyConsideringListOfKeys($expected, $actual, $keysToBeConsidered, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertArrayIsEqualToArrayIgnoringListOfKeys')) {
    function assertArrayIsEqualToArrayIgnoringListOfKeys(array $expected, array $actual, array $keysToBeIgnored, string $message = ''): void
    {
        Assert::assertArrayIsEqualToArrayIgnoringListOfKeys($expected, $actual, $keysToBeIgnored, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertArrayIsIdenticalToArrayOnlyConsideringListOfKeys')) {
    function assertArrayIsIdenticalToArrayOnlyConsideringListOfKeys(array $expected, array $actual, array $keysToBeConsidered, string $message = ''): void
    {
        Assert::assertArrayIsIdenticalToArrayOnlyConsideringListOfKeys($expected, $actual, $keysToBeConsidered, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertArrayIsIdenticalToArrayIgnoringListOfKeys')) {
    function assertArrayIsIdenticalToArrayIgnoringListOfKeys(array $expected, array $actual, array $keysToBeIgnored, string $message = ''): void
    {
        Assert::assertArrayIsIdenticalToArrayIgnoringListOfKeys($expected, $actual, $keysToBeIgnored, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertArrayHasKey')) {
    function assertArrayHasKey(mixed $key, array|ArrayAccess $array, string $message = ''): void
    {
        Assert::assertArrayHasKey($key, $array, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertArrayNotHasKey')) {
    function assertArrayNotHasKey(mixed $key, array|ArrayAccess $array, string $message = ''): void
    {
        Assert::assertArrayNotHasKey($key, $array, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsList')) {
    function assertIsList(mixed $array, string $message = ''): void
    {
        Assert::assertIsList($array, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertArraysAreIdentical')) {
    function assertArraysAreIdentical(array $expected, array $actual, string $message = ''): void
    {
        Assert::assertArraysAreIdentical($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertArraysAreIdenticalIgnoringOrder')) {
    function assertArraysAreIdenticalIgnoringOrder(array $expected, array $actual, string $message = ''): void
    {
        Assert::assertArraysAreIdenticalIgnoringOrder($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertArraysHaveIdenticalValues')) {
    function assertArraysHaveIdenticalValues(array $expected, array $actual, string $message = ''): void
    {
        Assert::assertArraysHaveIdenticalValues($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertArraysHaveIdenticalValuesIgnoringOrder')) {
    function assertArraysHaveIdenticalValuesIgnoringOrder(array $expected, array $actual, string $message = ''): void
    {
        Assert::assertArraysHaveIdenticalValuesIgnoringOrder($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertArraysAreEqual')) {
    function assertArraysAreEqual(array $expected, array $actual, string $message = ''): void
    {
        Assert::assertArraysAreEqual($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertArraysAreEqualIgnoringOrder')) {
    function assertArraysAreEqualIgnoringOrder(array $expected, array $actual, string $message = ''): void
    {
        Assert::assertArraysAreEqualIgnoringOrder($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertArraysHaveEqualValues')) {
    function assertArraysHaveEqualValues(array $expected, array $actual, string $message = ''): void
    {
        Assert::assertArraysHaveEqualValues($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertArraysHaveEqualValuesIgnoringOrder')) {
    function assertArraysHaveEqualValuesIgnoringOrder(array $expected, array $actual, string $message = ''): void
    {
        Assert::assertArraysHaveEqualValuesIgnoringOrder($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContains')) {
    function assertContains(mixed $needle, iterable $haystack, string $message = ''): void
    {
        Assert::assertContains($needle, $haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsEquals')) {
    function assertContainsEquals(mixed $needle, iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsEquals($needle, $haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNotContains')) {
    function assertNotContains(mixed $needle, iterable $haystack, string $message = ''): void
    {
        Assert::assertNotContains($needle, $haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNotContainsEquals')) {
    function assertNotContainsEquals(mixed $needle, iterable $haystack, string $message = ''): void
    {
        Assert::assertNotContainsEquals($needle, $haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsOnlyArray')) {
    function assertContainsOnlyArray(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsOnlyArray($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsOnlyBool')) {
    function assertContainsOnlyBool(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsOnlyBool($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsOnlyCallable')) {
    function assertContainsOnlyCallable(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsOnlyCallable($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsOnlyFloat')) {
    function assertContainsOnlyFloat(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsOnlyFloat($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsOnlyInt')) {
    function assertContainsOnlyInt(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsOnlyInt($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsOnlyIterable')) {
    function assertContainsOnlyIterable(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsOnlyIterable($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsOnlyNull')) {
    function assertContainsOnlyNull(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsOnlyNull($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsOnlyNumeric')) {
    function assertContainsOnlyNumeric(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsOnlyNumeric($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsOnlyObject')) {
    function assertContainsOnlyObject(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsOnlyObject($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsOnlyResource')) {
    function assertContainsOnlyResource(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsOnlyResource($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsOnlyClosedResource')) {
    function assertContainsOnlyClosedResource(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsOnlyClosedResource($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsOnlyScalar')) {
    function assertContainsOnlyScalar(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsOnlyScalar($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsOnlyString')) {
    function assertContainsOnlyString(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsOnlyString($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsOnlyInstancesOf')) {
    function assertContainsOnlyInstancesOf(string $className, iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsOnlyInstancesOf($className, $haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsNotOnlyArray')) {
    function assertContainsNotOnlyArray(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsNotOnlyArray($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsNotOnlyBool')) {
    function assertContainsNotOnlyBool(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsNotOnlyBool($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsNotOnlyCallable')) {
    function assertContainsNotOnlyCallable(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsNotOnlyCallable($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsNotOnlyFloat')) {
    function assertContainsNotOnlyFloat(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsNotOnlyFloat($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsNotOnlyInt')) {
    function assertContainsNotOnlyInt(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsNotOnlyInt($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsNotOnlyIterable')) {
    function assertContainsNotOnlyIterable(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsNotOnlyIterable($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsNotOnlyNull')) {
    function assertContainsNotOnlyNull(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsNotOnlyNull($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsNotOnlyNumeric')) {
    function assertContainsNotOnlyNumeric(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsNotOnlyNumeric($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsNotOnlyObject')) {
    function assertContainsNotOnlyObject(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsNotOnlyObject($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsNotOnlyResource')) {
    function assertContainsNotOnlyResource(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsNotOnlyResource($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsNotOnlyClosedResource')) {
    function assertContainsNotOnlyClosedResource(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsNotOnlyClosedResource($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsNotOnlyScalar')) {
    function assertContainsNotOnlyScalar(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsNotOnlyScalar($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsNotOnlyString')) {
    function assertContainsNotOnlyString(iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsNotOnlyString($haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertContainsNotOnlyInstancesOf')) {
    function assertContainsNotOnlyInstancesOf(string $className, iterable $haystack, string $message = ''): void
    {
        Assert::assertContainsNotOnlyInstancesOf($className, $haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertCount')) {
    function assertCount(int $expectedCount, Countable|iterable $haystack, string $message = ''): void
    {
        Assert::assertCount($expectedCount, $haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNotCount')) {
    function assertNotCount(int $expectedCount, Countable|iterable $haystack, string $message = ''): void
    {
        Assert::assertNotCount($expectedCount, $haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertEquals')) {
    function assertEquals(mixed $expected, mixed $actual, string $message = ''): void
    {
        Assert::assertEquals($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertEqualsCanonicalizing')) {
    function assertEqualsCanonicalizing(mixed $expected, mixed $actual, string $message = ''): void
    {
        Assert::assertEqualsCanonicalizing($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertEqualsIgnoringCase')) {
    function assertEqualsIgnoringCase(mixed $expected, mixed $actual, string $message = ''): void
    {
        Assert::assertEqualsIgnoringCase($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertEqualsWithDelta')) {
    function assertEqualsWithDelta(mixed $expected, mixed $actual, float $delta, string $message = ''): void
    {
        Assert::assertEqualsWithDelta($expected, $actual, $delta, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNotEquals')) {
    function assertNotEquals(mixed $expected, mixed $actual, string $message = ''): void
    {
        Assert::assertNotEquals($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNotEqualsCanonicalizing')) {
    function assertNotEqualsCanonicalizing(mixed $expected, mixed $actual, string $message = ''): void
    {
        Assert::assertNotEqualsCanonicalizing($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNotEqualsIgnoringCase')) {
    function assertNotEqualsIgnoringCase(mixed $expected, mixed $actual, string $message = ''): void
    {
        Assert::assertNotEqualsIgnoringCase($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNotEqualsWithDelta')) {
    function assertNotEqualsWithDelta(mixed $expected, mixed $actual, float $delta, string $message = ''): void
    {
        Assert::assertNotEqualsWithDelta($expected, $actual, $delta, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertObjectEquals')) {
    function assertObjectEquals(object $expected, object $actual, string $method = 'equals', string $message = ''): void
    {
        Assert::assertObjectEquals($expected, $actual, $method, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertObjectNotEquals')) {
    function assertObjectNotEquals(object $expected, object $actual, string $method = 'equals', string $message = ''): void
    {
        Assert::assertObjectNotEquals($expected, $actual, $method, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertEmpty')) {
    function assertEmpty(mixed $actual, string $message = ''): void
    {
        Assert::assertEmpty($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNotEmpty')) {
    function assertNotEmpty(mixed $actual, string $message = ''): void
    {
        Assert::assertNotEmpty($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertGreaterThan')) {
    function assertGreaterThan(mixed $minimum, mixed $actual, string $message = ''): void
    {
        Assert::assertGreaterThan($minimum, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertGreaterThanOrEqual')) {
    function assertGreaterThanOrEqual(mixed $minimum, mixed $actual, string $message = ''): void
    {
        Assert::assertGreaterThanOrEqual($minimum, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertLessThan')) {
    function assertLessThan(mixed $maximum, mixed $actual, string $message = ''): void
    {
        Assert::assertLessThan($maximum, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertLessThanOrEqual')) {
    function assertLessThanOrEqual(mixed $maximum, mixed $actual, string $message = ''): void
    {
        Assert::assertLessThanOrEqual($maximum, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileEquals')) {
    function assertFileEquals(string $expected, string $actual, string $message = ''): void
    {
        Assert::assertFileEquals($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileEqualsCanonicalizing')) {
    function assertFileEqualsCanonicalizing(string $expected, string $actual, string $message = ''): void
    {
        Assert::assertFileEqualsCanonicalizing($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileEqualsIgnoringCase')) {
    function assertFileEqualsIgnoringCase(string $expected, string $actual, string $message = ''): void
    {
        Assert::assertFileEqualsIgnoringCase($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileEqualsFileIgnoringWhitespace')) {
    function assertFileEqualsFileIgnoringWhitespace(string $expected, string $actual, string $message = ''): void
    {
        Assert::assertFileEqualsFileIgnoringWhitespace($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileNotEquals')) {
    function assertFileNotEquals(string $expected, string $actual, string $message = ''): void
    {
        Assert::assertFileNotEquals($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileNotEqualsCanonicalizing')) {
    function assertFileNotEqualsCanonicalizing(string $expected, string $actual, string $message = ''): void
    {
        Assert::assertFileNotEqualsCanonicalizing($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileNotEqualsIgnoringCase')) {
    function assertFileNotEqualsIgnoringCase(string $expected, string $actual, string $message = ''): void
    {
        Assert::assertFileNotEqualsIgnoringCase($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileNotEqualsFileIgnoringWhitespace')) {
    function assertFileNotEqualsFileIgnoringWhitespace(string $expected, string $actual, string $message = ''): void
    {
        Assert::assertFileNotEqualsFileIgnoringWhitespace($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringEqualsFile')) {
    function assertStringEqualsFile(string $expectedFile, string $actualString, string $message = ''): void
    {
        Assert::assertStringEqualsFile($expectedFile, $actualString, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringEqualsFileCanonicalizing')) {
    function assertStringEqualsFileCanonicalizing(string $expectedFile, string $actualString, string $message = ''): void
    {
        Assert::assertStringEqualsFileCanonicalizing($expectedFile, $actualString, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringEqualsFileIgnoringCase')) {
    function assertStringEqualsFileIgnoringCase(string $expectedFile, string $actualString, string $message = ''): void
    {
        Assert::assertStringEqualsFileIgnoringCase($expectedFile, $actualString, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringNotEqualsFile')) {
    function assertStringNotEqualsFile(string $expectedFile, string $actualString, string $message = ''): void
    {
        Assert::assertStringNotEqualsFile($expectedFile, $actualString, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringNotEqualsFileCanonicalizing')) {
    function assertStringNotEqualsFileCanonicalizing(string $expectedFile, string $actualString, string $message = ''): void
    {
        Assert::assertStringNotEqualsFileCanonicalizing($expectedFile, $actualString, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringNotEqualsFileIgnoringCase')) {
    function assertStringNotEqualsFileIgnoringCase(string $expectedFile, string $actualString, string $message = ''): void
    {
        Assert::assertStringNotEqualsFileIgnoringCase($expectedFile, $actualString, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringEqualsFileIgnoringWhitespace')) {
    function assertStringEqualsFileIgnoringWhitespace(string $expectedFile, string $actualString, string $message = ''): void
    {
        Assert::assertStringEqualsFileIgnoringWhitespace($expectedFile, $actualString, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringNotEqualsFileIgnoringWhitespace')) {
    function assertStringNotEqualsFileIgnoringWhitespace(string $expectedFile, string $actualString, string $message = ''): void
    {
        Assert::assertStringNotEqualsFileIgnoringWhitespace($expectedFile, $actualString, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsReadable')) {
    function assertIsReadable(string $filename, string $message = ''): void
    {
        Assert::assertIsReadable($filename, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsNotReadable')) {
    function assertIsNotReadable(string $filename, string $message = ''): void
    {
        Assert::assertIsNotReadable($filename, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsWritable')) {
    function assertIsWritable(string $filename, string $message = ''): void
    {
        Assert::assertIsWritable($filename, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsNotWritable')) {
    function assertIsNotWritable(string $filename, string $message = ''): void
    {
        Assert::assertIsNotWritable($filename, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertDirectoryExists')) {
    function assertDirectoryExists(string $directory, string $message = ''): void
    {
        Assert::assertDirectoryExists($directory, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertDirectoryDoesNotExist')) {
    function assertDirectoryDoesNotExist(string $directory, string $message = ''): void
    {
        Assert::assertDirectoryDoesNotExist($directory, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertDirectoryIsReadable')) {
    function assertDirectoryIsReadable(string $directory, string $message = ''): void
    {
        Assert::assertDirectoryIsReadable($directory, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertDirectoryIsNotReadable')) {
    function assertDirectoryIsNotReadable(string $directory, string $message = ''): void
    {
        Assert::assertDirectoryIsNotReadable($directory, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertDirectoryIsWritable')) {
    function assertDirectoryIsWritable(string $directory, string $message = ''): void
    {
        Assert::assertDirectoryIsWritable($directory, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertDirectoryIsNotWritable')) {
    function assertDirectoryIsNotWritable(string $directory, string $message = ''): void
    {
        Assert::assertDirectoryIsNotWritable($directory, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileExists')) {
    function assertFileExists(string $filename, string $message = ''): void
    {
        Assert::assertFileExists($filename, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileDoesNotExist')) {
    function assertFileDoesNotExist(string $filename, string $message = ''): void
    {
        Assert::assertFileDoesNotExist($filename, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileIsReadable')) {
    function assertFileIsReadable(string $file, string $message = ''): void
    {
        Assert::assertFileIsReadable($file, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileIsNotReadable')) {
    function assertFileIsNotReadable(string $file, string $message = ''): void
    {
        Assert::assertFileIsNotReadable($file, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileIsWritable')) {
    function assertFileIsWritable(string $file, string $message = ''): void
    {
        Assert::assertFileIsWritable($file, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileIsNotWritable')) {
    function assertFileIsNotWritable(string $file, string $message = ''): void
    {
        Assert::assertFileIsNotWritable($file, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertTrue')) {
    function assertTrue(mixed $condition, string $message = ''): void
    {
        Assert::assertTrue($condition, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNotTrue')) {
    function assertNotTrue(mixed $condition, string $message = ''): void
    {
        Assert::assertNotTrue($condition, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFalse')) {
    function assertFalse(mixed $condition, string $message = ''): void
    {
        Assert::assertFalse($condition, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNotFalse')) {
    function assertNotFalse(mixed $condition, string $message = ''): void
    {
        Assert::assertNotFalse($condition, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNull')) {
    function assertNull(mixed $actual, string $message = ''): void
    {
        Assert::assertNull($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNotNull')) {
    function assertNotNull(mixed $actual, string $message = ''): void
    {
        Assert::assertNotNull($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFinite')) {
    function assertFinite(mixed $actual, string $message = ''): void
    {
        Assert::assertFinite($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertInfinite')) {
    function assertInfinite(mixed $actual, string $message = ''): void
    {
        Assert::assertInfinite($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNan')) {
    function assertNan(mixed $actual, string $message = ''): void
    {
        Assert::assertNan($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertObjectHasProperty')) {
    function assertObjectHasProperty(string $propertyName, object $object, string $message = ''): void
    {
        Assert::assertObjectHasProperty($propertyName, $object, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertObjectNotHasProperty')) {
    function assertObjectNotHasProperty(string $propertyName, object $object, string $message = ''): void
    {
        Assert::assertObjectNotHasProperty($propertyName, $object, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertSame')) {
    function assertSame(mixed $expected, mixed $actual, string $message = ''): void
    {
        Assert::assertSame($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNotSame')) {
    function assertNotSame(mixed $expected, mixed $actual, string $message = ''): void
    {
        Assert::assertNotSame($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertInstanceOf')) {
    function assertInstanceOf(string $expected, mixed $actual, string $message = ''): void
    {
        Assert::assertInstanceOf($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNotInstanceOf')) {
    function assertNotInstanceOf(string $expected, mixed $actual, string $message = ''): void
    {
        Assert::assertNotInstanceOf($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsArray')) {
    function assertIsArray(mixed $actual, string $message = ''): void
    {
        Assert::assertIsArray($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsBool')) {
    function assertIsBool(mixed $actual, string $message = ''): void
    {
        Assert::assertIsBool($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsFloat')) {
    function assertIsFloat(mixed $actual, string $message = ''): void
    {
        Assert::assertIsFloat($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsInt')) {
    function assertIsInt(mixed $actual, string $message = ''): void
    {
        Assert::assertIsInt($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsNumeric')) {
    function assertIsNumeric(mixed $actual, string $message = ''): void
    {
        Assert::assertIsNumeric($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsObject')) {
    function assertIsObject(mixed $actual, string $message = ''): void
    {
        Assert::assertIsObject($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsResource')) {
    function assertIsResource(mixed $actual, string $message = ''): void
    {
        Assert::assertIsResource($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsClosedResource')) {
    function assertIsClosedResource(mixed $actual, string $message = ''): void
    {
        Assert::assertIsClosedResource($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsString')) {
    function assertIsString(mixed $actual, string $message = ''): void
    {
        Assert::assertIsString($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsScalar')) {
    function assertIsScalar(mixed $actual, string $message = ''): void
    {
        Assert::assertIsScalar($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsCallable')) {
    function assertIsCallable(mixed $actual, string $message = ''): void
    {
        Assert::assertIsCallable($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsIterable')) {
    function assertIsIterable(mixed $actual, string $message = ''): void
    {
        Assert::assertIsIterable($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsNotArray')) {
    function assertIsNotArray(mixed $actual, string $message = ''): void
    {
        Assert::assertIsNotArray($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsNotBool')) {
    function assertIsNotBool(mixed $actual, string $message = ''): void
    {
        Assert::assertIsNotBool($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsNotFloat')) {
    function assertIsNotFloat(mixed $actual, string $message = ''): void
    {
        Assert::assertIsNotFloat($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsNotInt')) {
    function assertIsNotInt(mixed $actual, string $message = ''): void
    {
        Assert::assertIsNotInt($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsNotNumeric')) {
    function assertIsNotNumeric(mixed $actual, string $message = ''): void
    {
        Assert::assertIsNotNumeric($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsNotObject')) {
    function assertIsNotObject(mixed $actual, string $message = ''): void
    {
        Assert::assertIsNotObject($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsNotResource')) {
    function assertIsNotResource(mixed $actual, string $message = ''): void
    {
        Assert::assertIsNotResource($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsNotClosedResource')) {
    function assertIsNotClosedResource(mixed $actual, string $message = ''): void
    {
        Assert::assertIsNotClosedResource($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsNotString')) {
    function assertIsNotString(mixed $actual, string $message = ''): void
    {
        Assert::assertIsNotString($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsNotScalar')) {
    function assertIsNotScalar(mixed $actual, string $message = ''): void
    {
        Assert::assertIsNotScalar($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsNotCallable')) {
    function assertIsNotCallable(mixed $actual, string $message = ''): void
    {
        Assert::assertIsNotCallable($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertIsNotIterable')) {
    function assertIsNotIterable(mixed $actual, string $message = ''): void
    {
        Assert::assertIsNotIterable($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertMatchesRegularExpression')) {
    function assertMatchesRegularExpression(string $pattern, string $string, string $message = ''): void
    {
        Assert::assertMatchesRegularExpression($pattern, $string, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertDoesNotMatchRegularExpression')) {
    function assertDoesNotMatchRegularExpression(string $pattern, string $string, string $message = ''): void
    {
        Assert::assertDoesNotMatchRegularExpression($pattern, $string, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertSameSize')) {
    function assertSameSize(Countable|iterable $expected, Countable|iterable $actual, string $message = ''): void
    {
        Assert::assertSameSize($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertNotSameSize')) {
    function assertNotSameSize(Countable|iterable $expected, Countable|iterable $actual, string $message = ''): void
    {
        Assert::assertNotSameSize($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringContainsStringIgnoringLineEndings')) {
    function assertStringContainsStringIgnoringLineEndings(string $needle, string $haystack, string $message = ''): void
    {
        Assert::assertStringContainsStringIgnoringLineEndings($needle, $haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringEqualsStringIgnoringLineEndings')) {
    function assertStringEqualsStringIgnoringLineEndings(string $expected, string $actual, string $message = ''): void
    {
        Assert::assertStringEqualsStringIgnoringLineEndings($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringEqualsStringIgnoringWhitespace')) {
    function assertStringEqualsStringIgnoringWhitespace(string $expected, string $actual, string $message = ''): void
    {
        Assert::assertStringEqualsStringIgnoringWhitespace($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringNotEqualsStringIgnoringWhitespace')) {
    function assertStringNotEqualsStringIgnoringWhitespace(string $expected, string $actual, string $message = ''): void
    {
        Assert::assertStringNotEqualsStringIgnoringWhitespace($expected, $actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileMatchesFormat')) {
    function assertFileMatchesFormat(string $format, string $actualFile, string $message = ''): void
    {
        Assert::assertFileMatchesFormat($format, $actualFile, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertFileMatchesFormatFile')) {
    function assertFileMatchesFormatFile(string $formatFile, string $actualFile, string $message = ''): void
    {
        Assert::assertFileMatchesFormatFile($formatFile, $actualFile, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringMatchesFormat')) {
    function assertStringMatchesFormat(string $format, string $string, string $message = ''): void
    {
        Assert::assertStringMatchesFormat($format, $string, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringMatchesFormatFile')) {
    function assertStringMatchesFormatFile(string $formatFile, string $string, string $message = ''): void
    {
        Assert::assertStringMatchesFormatFile($formatFile, $string, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringStartsWith')) {
    function assertStringStartsWith(string $prefix, string $string, string $message = ''): void
    {
        Assert::assertStringStartsWith($prefix, $string, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringStartsNotWith')) {
    function assertStringStartsNotWith(string $prefix, string $string, string $message = ''): void
    {
        Assert::assertStringStartsNotWith($prefix, $string, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringContainsString')) {
    function assertStringContainsString(string $needle, string $haystack, string $message = ''): void
    {
        Assert::assertStringContainsString($needle, $haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringContainsStringIgnoringCase')) {
    function assertStringContainsStringIgnoringCase(string $needle, string $haystack, string $message = ''): void
    {
        Assert::assertStringContainsStringIgnoringCase($needle, $haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringNotContainsString')) {
    function assertStringNotContainsString(string $needle, string $haystack, string $message = ''): void
    {
        Assert::assertStringNotContainsString($needle, $haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringNotContainsStringIgnoringCase')) {
    function assertStringNotContainsStringIgnoringCase(string $needle, string $haystack, string $message = ''): void
    {
        Assert::assertStringNotContainsStringIgnoringCase($needle, $haystack, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringEndsWith')) {
    function assertStringEndsWith(string $suffix, string $string, string $message = ''): void
    {
        Assert::assertStringEndsWith($suffix, $string, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertStringEndsNotWith')) {
    function assertStringEndsNotWith(string $suffix, string $string, string $message = ''): void
    {
        Assert::assertStringEndsNotWith($suffix, $string, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertXmlFileEqualsXmlFile')) {
    function assertXmlFileEqualsXmlFile(string $expectedFile, string $actualFile, string $message = ''): void
    {
        Assert::assertXmlFileEqualsXmlFile($expectedFile, $actualFile, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertXmlFileNotEqualsXmlFile')) {
    function assertXmlFileNotEqualsXmlFile(string $expectedFile, string $actualFile, string $message = ''): void
    {
        Assert::assertXmlFileNotEqualsXmlFile($expectedFile, $actualFile, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertXmlStringEqualsXmlFile')) {
    function assertXmlStringEqualsXmlFile(string $expectedFile, string $actualXml, string $message = ''): void
    {
        Assert::assertXmlStringEqualsXmlFile($expectedFile, $actualXml, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertXmlStringNotEqualsXmlFile')) {
    function assertXmlStringNotEqualsXmlFile(string $expectedFile, string $actualXml, string $message = ''): void
    {
        Assert::assertXmlStringNotEqualsXmlFile($expectedFile, $actualXml, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertXmlStringEqualsXmlString')) {
    function assertXmlStringEqualsXmlString(string $expectedXml, string $actualXml, string $message = ''): void
    {
        Assert::assertXmlStringEqualsXmlString($expectedXml, $actualXml, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertXmlStringNotEqualsXmlString')) {
    function assertXmlStringNotEqualsXmlString(string $expectedXml, string $actualXml, string $message = ''): void
    {
        Assert::assertXmlStringNotEqualsXmlString($expectedXml, $actualXml, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertXmlFileEqualsXmlFileConsideringComments')) {
    function assertXmlFileEqualsXmlFileConsideringComments(string $expectedFile, string $actualFile, string $message = ''): void
    {
        Assert::assertXmlFileEqualsXmlFileConsideringComments($expectedFile, $actualFile, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertXmlFileNotEqualsXmlFileConsideringComments')) {
    function assertXmlFileNotEqualsXmlFileConsideringComments(string $expectedFile, string $actualFile, string $message = ''): void
    {
        Assert::assertXmlFileNotEqualsXmlFileConsideringComments($expectedFile, $actualFile, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertXmlStringEqualsXmlFileConsideringComments')) {
    function assertXmlStringEqualsXmlFileConsideringComments(string $expectedFile, string $actualXml, string $message = ''): void
    {
        Assert::assertXmlStringEqualsXmlFileConsideringComments($expectedFile, $actualXml, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertXmlStringNotEqualsXmlFileConsideringComments')) {
    function assertXmlStringNotEqualsXmlFileConsideringComments(string $expectedFile, string $actualXml, string $message = ''): void
    {
        Assert::assertXmlStringNotEqualsXmlFileConsideringComments($expectedFile, $actualXml, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertXmlStringEqualsXmlStringConsideringComments')) {
    function assertXmlStringEqualsXmlStringConsideringComments(string $expectedXml, string $actualXml, string $message = ''): void
    {
        Assert::assertXmlStringEqualsXmlStringConsideringComments($expectedXml, $actualXml, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertXmlStringNotEqualsXmlStringConsideringComments')) {
    function assertXmlStringNotEqualsXmlStringConsideringComments(string $expectedXml, string $actualXml, string $message = ''): void
    {
        Assert::assertXmlStringNotEqualsXmlStringConsideringComments($expectedXml, $actualXml, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertThat')) {
    function assertThat(mixed $value, Constraint $constraint, string $message = ''): void
    {
        Assert::assertThat($value, $constraint, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertJson')) {
    function assertJson(string $actual, string $message = ''): void
    {
        Assert::assertJson($actual, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertJsonStringEqualsJsonString')) {
    function assertJsonStringEqualsJsonString(string $expectedJson, string $actualJson, string $message = ''): void
    {
        Assert::assertJsonStringEqualsJsonString($expectedJson, $actualJson, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertJsonStringNotEqualsJsonString')) {
    function assertJsonStringNotEqualsJsonString(string $expectedJson, string $actualJson, string $message = ''): void
    {
        Assert::assertJsonStringNotEqualsJsonString($expectedJson, $actualJson, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertJsonStringEqualsJsonFile')) {
    function assertJsonStringEqualsJsonFile(string $expectedFile, string $actualJson, string $message = ''): void
    {
        Assert::assertJsonStringEqualsJsonFile($expectedFile, $actualJson, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertJsonStringNotEqualsJsonFile')) {
    function assertJsonStringNotEqualsJsonFile(string $expectedFile, string $actualJson, string $message = ''): void
    {
        Assert::assertJsonStringNotEqualsJsonFile($expectedFile, $actualJson, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertJsonFileEqualsJsonFile')) {
    function assertJsonFileEqualsJsonFile(string $expectedFile, string $actualFile, string $message = ''): void
    {
        Assert::assertJsonFileEqualsJsonFile($expectedFile, $actualFile, $message);
    }
}
if (!function_exists('PHPUnit\Framework\assertJsonFileNotEqualsJsonFile')) {
    function assertJsonFileNotEqualsJsonFile(string $expectedFile, string $actualFile, string $message = ''): void
    {
        Assert::assertJsonFileNotEqualsJsonFile($expectedFile, $actualFile, $message);
    }
}
if (!function_exists('PHPUnit\Framework\logicalAnd')) {
    function logicalAnd(mixed ...$constraints): LogicalAnd
    {
        return Assert::logicalAnd(...$constraints);
    }
}
if (!function_exists('PHPUnit\Framework\logicalOr')) {
    function logicalOr(mixed ...$constraints): LogicalOr
    {
        return Assert::logicalOr(...$constraints);
    }
}
if (!function_exists('PHPUnit\Framework\logicalNot')) {
    function logicalNot(Constraint $constraint): LogicalNot
    {
        return Assert::logicalNot($constraint);
    }
}
if (!function_exists('PHPUnit\Framework\logicalXor')) {
    function logicalXor(mixed ...$constraints): LogicalXor
    {
        return Assert::logicalXor(...$constraints);
    }
}
if (!function_exists('PHPUnit\Framework\anything')) {
    function anything(): IsAnything
    {
        return Assert::anything();
    }
}
if (!function_exists('PHPUnit\Framework\isTrue')) {
    function isTrue(): IsTrue
    {
        return Assert::isTrue();
    }
}
if (!function_exists('PHPUnit\Framework\isFalse')) {
    function isFalse(): IsFalse
    {
        return Assert::isFalse();
    }
}
if (!function_exists('PHPUnit\Framework\isJson')) {
    function isJson(): IsJson
    {
        return Assert::isJson();
    }
}
if (!function_exists('PHPUnit\Framework\isNull')) {
    function isNull(): IsNull
    {
        return Assert::isNull();
    }
}
if (!function_exists('PHPUnit\Framework\isFinite')) {
    function isFinite(): IsFinite
    {
        return Assert::isFinite();
    }
}
if (!function_exists('PHPUnit\Framework\isInfinite')) {
    function isInfinite(): IsInfinite
    {
        return Assert::isInfinite();
    }
}
if (!function_exists('PHPUnit\Framework\isNan')) {
    function isNan(): IsNan
    {
        return Assert::isNan();
    }
}
if (!function_exists('PHPUnit\Framework\containsEqual')) {
    function containsEqual(mixed $value): TraversableContainsEqual
    {
        return Assert::containsEqual($value);
    }
}
if (!function_exists('PHPUnit\Framework\containsIdentical')) {
    function containsIdentical(mixed $value): TraversableContainsIdentical
    {
        return Assert::containsIdentical($value);
    }
}
if (!function_exists('PHPUnit\Framework\containsOnlyArray')) {
    function containsOnlyArray(): TraversableContainsOnly
    {
        return Assert::containsOnlyArray();
    }
}
if (!function_exists('PHPUnit\Framework\containsOnlyBool')) {
    function containsOnlyBool(): TraversableContainsOnly
    {
        return Assert::containsOnlyBool();
    }
}
if (!function_exists('PHPUnit\Framework\containsOnlyCallable')) {
    function containsOnlyCallable(): TraversableContainsOnly
    {
        return Assert::containsOnlyCallable();
    }
}
if (!function_exists('PHPUnit\Framework\containsOnlyFloat')) {
    function containsOnlyFloat(): TraversableContainsOnly
    {
        return Assert::containsOnlyFloat();
    }
}
if (!function_exists('PHPUnit\Framework\containsOnlyInt')) {
    function containsOnlyInt(): TraversableContainsOnly
    {
        return Assert::containsOnlyInt();
    }
}
if (!function_exists('PHPUnit\Framework\containsOnlyIterable')) {
    function containsOnlyIterable(): TraversableContainsOnly
    {
        return Assert::containsOnlyIterable();
    }
}
if (!function_exists('PHPUnit\Framework\containsOnlyNull')) {
    function containsOnlyNull(): TraversableContainsOnly
    {
        return Assert::containsOnlyNull();
    }
}
if (!function_exists('PHPUnit\Framework\containsOnlyNumeric')) {
    function containsOnlyNumeric(): TraversableContainsOnly
    {
        return Assert::containsOnlyNumeric();
    }
}
if (!function_exists('PHPUnit\Framework\containsOnlyObject')) {
    function containsOnlyObject(): TraversableContainsOnly
    {
        return Assert::containsOnlyObject();
    }
}
if (!function_exists('PHPUnit\Framework\containsOnlyResource')) {
    function containsOnlyResource(): TraversableContainsOnly
    {
        return Assert::containsOnlyResource();
    }
}
if (!function_exists('PHPUnit\Framework\containsOnlyClosedResource')) {
    function containsOnlyClosedResource(): TraversableContainsOnly
    {
        return Assert::containsOnlyClosedResource();
    }
}
if (!function_exists('PHPUnit\Framework\containsOnlyScalar')) {
    function containsOnlyScalar(): TraversableContainsOnly
    {
        return Assert::containsOnlyScalar();
    }
}
if (!function_exists('PHPUnit\Framework\containsOnlyString')) {
    function containsOnlyString(): TraversableContainsOnly
    {
        return Assert::containsOnlyString();
    }
}
if (!function_exists('PHPUnit\Framework\containsOnlyInstancesOf')) {
    function containsOnlyInstancesOf(string $className): TraversableContainsOnly
    {
        return Assert::containsOnlyInstancesOf($className);
    }
}
if (!function_exists('PHPUnit\Framework\arrayHasKey')) {
    function arrayHasKey(mixed $key): ArrayHasKey
    {
        return Assert::arrayHasKey($key);
    }
}
if (!function_exists('PHPUnit\Framework\isList')) {
    function isList(): IsList
    {
        return Assert::isList();
    }
}
if (!function_exists('PHPUnit\Framework\equalTo')) {
    function equalTo(mixed $value): IsEqual
    {
        return Assert::equalTo($value);
    }
}
if (!function_exists('PHPUnit\Framework\equalToCanonicalizing')) {
    function equalToCanonicalizing(mixed $value): IsEqualCanonicalizing
    {
        return Assert::equalToCanonicalizing($value);
    }
}
if (!function_exists('PHPUnit\Framework\equalToIgnoringCase')) {
    function equalToIgnoringCase(mixed $value): IsEqualIgnoringCase
    {
        return Assert::equalToIgnoringCase($value);
    }
}
if (!function_exists('PHPUnit\Framework\equalToWithDelta')) {
    function equalToWithDelta(mixed $value, float $delta): IsEqualWithDelta
    {
        return Assert::equalToWithDelta($value, $delta);
    }
}
if (!function_exists('PHPUnit\Framework\isEmpty')) {
    function isEmpty(): IsEmpty
    {
        return Assert::isEmpty();
    }
}
if (!function_exists('PHPUnit\Framework\isWritable')) {
    function isWritable(): IsWritable
    {
        return Assert::isWritable();
    }
}
if (!function_exists('PHPUnit\Framework\isReadable')) {
    function isReadable(): IsReadable
    {
        return Assert::isReadable();
    }
}
if (!function_exists('PHPUnit\Framework\directoryExists')) {
    function directoryExists(): DirectoryExists
    {
        return Assert::directoryExists();
    }
}
if (!function_exists('PHPUnit\Framework\fileExists')) {
    function fileExists(): FileExists
    {
        return Assert::fileExists();
    }
}
if (!function_exists('PHPUnit\Framework\greaterThan')) {
    function greaterThan(mixed $value): GreaterThan
    {
        return Assert::greaterThan($value);
    }
}
if (!function_exists('PHPUnit\Framework\greaterThanOrEqual')) {
    function greaterThanOrEqual(mixed $value): LogicalOr
    {
        return Assert::greaterThanOrEqual($value);
    }
}
if (!function_exists('PHPUnit\Framework\identicalTo')) {
    function identicalTo(mixed $value): IsIdentical
    {
        return Assert::identicalTo($value);
    }
}
if (!function_exists('PHPUnit\Framework\isInstanceOf')) {
    function isInstanceOf(string $className): IsInstanceOf
    {
        return Assert::isInstanceOf($className);
    }
}
if (!function_exists('PHPUnit\Framework\isArray')) {
    function isArray(): IsType
    {
        return Assert::isArray();
    }
}
if (!function_exists('PHPUnit\Framework\isBool')) {
    function isBool(): IsType
    {
        return Assert::isBool();
    }
}
if (!function_exists('PHPUnit\Framework\isCallable')) {
    function isCallable(): IsType
    {
        return Assert::isCallable();
    }
}
if (!function_exists('PHPUnit\Framework\isFloat')) {
    function isFloat(): IsType
    {
        return Assert::isFloat();
    }
}
if (!function_exists('PHPUnit\Framework\isInt')) {
    function isInt(): IsType
    {
        return Assert::isInt();
    }
}
if (!function_exists('PHPUnit\Framework\isIterable')) {
    function isIterable(): IsType
    {
        return Assert::isIterable();
    }
}
if (!function_exists('PHPUnit\Framework\isNumeric')) {
    function isNumeric(): IsType
    {
        return Assert::isNumeric();
    }
}
if (!function_exists('PHPUnit\Framework\isObject')) {
    function isObject(): IsType
    {
        return Assert::isObject();
    }
}
if (!function_exists('PHPUnit\Framework\isResource')) {
    function isResource(): IsType
    {
        return Assert::isResource();
    }
}
if (!function_exists('PHPUnit\Framework\isClosedResource')) {
    function isClosedResource(): IsType
    {
        return Assert::isClosedResource();
    }
}
if (!function_exists('PHPUnit\Framework\isScalar')) {
    function isScalar(): IsType
    {
        return Assert::isScalar();
    }
}
if (!function_exists('PHPUnit\Framework\isString')) {
    function isString(): IsType
    {
        return Assert::isString();
    }
}
if (!function_exists('PHPUnit\Framework\lessThan')) {
    function lessThan(mixed $value): LessThan
    {
        return Assert::lessThan($value);
    }
}
if (!function_exists('PHPUnit\Framework\lessThanOrEqual')) {
    function lessThanOrEqual(mixed $value): LogicalOr
    {
        return Assert::lessThanOrEqual($value);
    }
}
if (!function_exists('PHPUnit\Framework\matchesRegularExpression')) {
    function matchesRegularExpression(string $pattern): RegularExpression
    {
        return Assert::matchesRegularExpression($pattern);
    }
}
if (!function_exists('PHPUnit\Framework\matches')) {
    function matches(string $string): StringMatchesFormatDescription
    {
        return Assert::matches($string);
    }
}
if (!function_exists('PHPUnit\Framework\stringStartsWith')) {
    function stringStartsWith(string $prefix): StringStartsWith
    {
        return Assert::stringStartsWith($prefix);
    }
}
if (!function_exists('PHPUnit\Framework\stringContains')) {
    function stringContains(string $string, bool $case = true): StringContains
    {
        return Assert::stringContains($string, $case);
    }
}
if (!function_exists('PHPUnit\Framework\stringEndsWith')) {
    function stringEndsWith(string $suffix): StringEndsWith
    {
        return Assert::stringEndsWith($suffix);
    }
}
if (!function_exists('PHPUnit\Framework\stringEqualsStringIgnoringLineEndings')) {
    function stringEqualsStringIgnoringLineEndings(string $string): StringEqualsStringIgnoringLineEndings
    {
        return Assert::stringEqualsStringIgnoringLineEndings($string);
    }
}
if (!function_exists('PHPUnit\Framework\stringEqualsStringIgnoringWhitespace')) {
    function stringEqualsStringIgnoringWhitespace(string $string): StringEqualsStringIgnoringWhitespace
    {
        return Assert::stringEqualsStringIgnoringWhitespace($string);
    }
}
if (!function_exists('PHPUnit\Framework\countOf')) {
    function countOf(int $count): Count
    {
        return Assert::countOf($count);
    }
}
if (!function_exists('PHPUnit\Framework\objectEquals')) {
    function objectEquals(object $object, string $method = 'equals'): ObjectEquals
    {
        return Assert::objectEquals($object, $method);
    }
}
if (!function_exists('PHPUnit\Framework\callback')) {
    function callback(callable $callback): Callback
    {
        return Assert::callback($callback);
    }
}
if (!function_exists('PHPUnit\Framework\any')) {
    function any(): AnyInvokedCountMatcher
    {
        return new AnyInvokedCountMatcher;
    }
}
if (!function_exists('PHPUnit\Framework\never')) {
    function never(): InvokedCountMatcher
    {
        return new InvokedCountMatcher(0);
    }
}
if (!function_exists('PHPUnit\Framework\atLeast')) {
    function atLeast(int $requiredInvocations): InvokedAtLeastCountMatcher
    {
        return new InvokedAtLeastCountMatcher(
            $requiredInvocations,
        );
    }
}
if (!function_exists('PHPUnit\Framework\atLeastOnce')) {
    function atLeastOnce(): InvokedAtLeastOnceMatcher
    {
        return new InvokedAtLeastOnceMatcher;
    }
}
if (!function_exists('PHPUnit\Framework\once')) {
    function once(): InvokedCountMatcher
    {
        return new InvokedCountMatcher(1);
    }
}
if (!function_exists('PHPUnit\Framework\exactly')) {
    function exactly(int $count): InvokedCountMatcher
    {
        return new InvokedCountMatcher($count);
    }
}
if (!function_exists('PHPUnit\Framework\atMost')) {
    function atMost(int $allowedInvocations): InvokedAtMostCountMatcher
    {
        return new InvokedAtMostCountMatcher($allowedInvocations);
    }
}
if (!function_exists('PHPUnit\Framework\throwException')) {
    function throwException(Throwable $exception): ExceptionStub
    {
        return new ExceptionStub($exception);
    }
}
