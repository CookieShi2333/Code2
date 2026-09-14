package org.apache.commons.lang3;
import java.lang.reflect.Array;
import java.lang.reflect.Field;
import java.lang.reflect.Method;
import java.lang.reflect.Type;
import java.security.SecureRandom;
import java.util.Arrays;
import java.util.BitSet;
import java.util.Comparator;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import java.util.Random;
import java.util.concurrent.ThreadLocalRandom;
import java.util.function.Function;
import java.util.function.IntFunction;
import java.util.function.Supplier;
import org.apache.commons.lang3.builder.EqualsBuilder;
import org.apache.commons.lang3.builder.HashCodeBuilder;
import org.apache.commons.lang3.builder.ToStringBuilder;
import org.apache.commons.lang3.builder.ToStringStyle;
import org.apache.commons.lang3.function.FailableFunction;
import org.apache.commons.lang3.mutable.MutableInt;
import org.apache.commons.lang3.stream.IntStreams;
import org.apache.commons.lang3.stream.Streams;
public class ArrayUtils {
    static class MathBridge {
        static int addExact(final int a, final int b) {
            return Math.addExact(a, b);
        }
    }
    public static final boolean[] EMPTY_BOOLEAN_ARRAY = {};
    public static final Boolean[] EMPTY_BOOLEAN_OBJECT_ARRAY = {};
    public static final byte[] EMPTY_BYTE_ARRAY = {};
    public static final Byte[] EMPTY_BYTE_OBJECT_ARRAY = {};
    public static final char[] EMPTY_CHAR_ARRAY = {};
    public static final Character[] EMPTY_CHARACTER_OBJECT_ARRAY = {};
    public static final Class<?>[] EMPTY_CLASS_ARRAY = {};
    public static final double[] EMPTY_DOUBLE_ARRAY = {};
    public static final Double[] EMPTY_DOUBLE_OBJECT_ARRAY = {};
    public static final Field[] EMPTY_FIELD_ARRAY = {};
    public static final float[] EMPTY_FLOAT_ARRAY = {};
    public static final Float[] EMPTY_FLOAT_OBJECT_ARRAY = {};
    public static final int[] EMPTY_INT_ARRAY = {};
    public static final Integer[] EMPTY_INTEGER_OBJECT_ARRAY = {};
    public static final long[] EMPTY_LONG_ARRAY = {};
    public static final Long[] EMPTY_LONG_OBJECT_ARRAY = {};
    public static final Method[] EMPTY_METHOD_ARRAY = {};
    public static final Object[] EMPTY_OBJECT_ARRAY = {};
    public static final short[] EMPTY_SHORT_ARRAY = {};
    public static final Short[] EMPTY_SHORT_OBJECT_ARRAY = {};
    public static final String[] EMPTY_STRING_ARRAY = {};
    public static final Throwable[] EMPTY_THROWABLE_ARRAY = {};
    public static final Type[] EMPTY_TYPE_ARRAY = {};
    public static final int INDEX_NOT_FOUND = -1;
    @Deprecated
    public static int SOFT_MAX_ARRAY_LENGTH = Integer.MAX_VALUE - 8;
    public static final int SAFE_MAX_ARRAY_LENGTH = Integer.MAX_VALUE - 8;
    public static boolean[] add(final boolean[] array, final boolean element) {
        final boolean[] newArray = (boolean[]) copyArrayGrow1(array, Boolean.TYPE);
        newArray[newArray.length - 1] = element;
        return newArray;
    }
    @Deprecated
    public static boolean[] add(final boolean[] array, final int index, final boolean element) {
        return (boolean[]) add(array, index, Boolean.valueOf(element), Boolean.TYPE);
    }
    public static byte[] add(final byte[] array, final byte element) {
        final byte[] newArray = (byte[]) copyArrayGrow1(array, Byte.TYPE);
        newArray[newArray.length - 1] = element;
        return newArray;
    }
    @Deprecated
    public static byte[] add(final byte[] array, final int index, final byte element) {
        return (byte[]) add(array, index, Byte.valueOf(element), Byte.TYPE);
    }
    public static char[] add(final char[] array, final char element) {
        final char[] newArray = (char[]) copyArrayGrow1(array, Character.TYPE);
        newArray[newArray.length - 1] = element;
        return newArray;
    }
    @Deprecated
    public static char[] add(final char[] array, final int index, final char element) {
        return (char[]) add(array, index, Character.valueOf(element), Character.TYPE);
    }
    public static double[] add(final double[] array, final double element) {
        final double[] newArray = (double[]) copyArrayGrow1(array, Double.TYPE);
        newArray[newArray.length - 1] = element;
        return newArray;
    }
    @Deprecated
    public static double[] add(final double[] array, final int index, final double element) {
        return (double[]) add(array, index, Double.valueOf(element), Double.TYPE);
    }
    public static float[] add(final float[] array, final float element) {
        final float[] newArray = (float[]) copyArrayGrow1(array, Float.TYPE);
        newArray[newArray.length - 1] = element;
        return newArray;
    }
    @Deprecated
    public static float[] add(final float[] array, final int index, final float element) {
        return (float[]) add(array, index, Float.valueOf(element), Float.TYPE);
    }
    public static int[] add(final int[] array, final int element) {
        final int[] newArray = (int[]) copyArrayGrow1(array, Integer.TYPE);
        newArray[newArray.length - 1] = element;
        return newArray;
    }
    @Deprecated
    public static int[] add(final int[] array, final int index, final int element) {
        return (int[]) add(array, index, Integer.valueOf(element), Integer.TYPE);
    }
    @Deprecated
    public static long[] add(final long[] array, final int index, final long element) {
        return (long[]) add(array, index, Long.valueOf(element), Long.TYPE);
    }
    public static long[] add(final long[] array, final long element) {
        final long[] newArray = (long[]) copyArrayGrow1(array, Long.TYPE);
        newArray[newArray.length - 1] = element;
        return newArray;
    }
    private static Object add(final Object array, final int index, final Object element, final Class<?> clazz) {
        if (array == null) {
            if (index != 0) {
                throw new IndexOutOfBoundsException("Index: " + index + ", Length: 0");
            }
            final Object joinedArray = Array.newInstance(clazz, 1);
            Array.set(joinedArray, 0, element);
            return joinedArray;
        }
        final int length = Array.getLength(array);
        if (index > length || index < 0) {
            throw new IndexOutOfBoundsException("Index: " + index + ", Length: " + length);
        }
        final Object result = arraycopy(array, 0, 0, index, () -> Array.newInstance(clazz, length + 1));
        Array.set(result, index, element);
        if (index < length) {
            System.arraycopy(array, index, result, index + 1, length - index);
        }
        return result;
    }
    @Deprecated
    public static short[] add(final short[] array, final int index, final short element) {
        return (short[]) add(array, index, Short.valueOf(element), Short.TYPE);
    }
    public static short[] add(final short[] array, final short element) {
        final short[] newArray = (short[]) copyArrayGrow1(array, Short.TYPE);
        newArray[newArray.length - 1] = element;
        return newArray;
    }
    @Deprecated
    public static <T> T[] add(final T[] array, final int index, final T element) {
        final Class<T> clazz;
        if (array != null) {
            clazz = getComponentType(array);
        } else if (element != null) {
            clazz = ObjectUtils.getClass(element);
        } else {
            throw new IllegalArgumentException("Array and element cannot both be null");
        }
        return (T[]) add(array, index, element, clazz);
    }
    public static <T> T[] add(final T[] array, final T element) {
        final Class<?> type;
        if (array != null) {
            type = array.getClass().getComponentType();
        } else if (element != null) {
            type = element.getClass();
        } else {
            throw new IllegalArgumentException("Arguments cannot both be null");
        }
        @SuppressWarnings("unchecked") 
        final
        T[] newArray = (T[]) copyArrayGrow1(array, type);
        newArray[newArray.length - 1] = element;
        return newArray;
    }
    public static boolean[] addAll(final boolean[] array1, final boolean... array2) {
        if (array1 == null) {
            return clone(array2);
        }
        if (array2 == null) {
            return clone(array1);
        }
        final boolean[] joinedArray = new boolean[addExact(array1.length, array2)];
        System.arraycopy(array1, 0, joinedArray, 0, array1.length);
        System.arraycopy(array2, 0, joinedArray, array1.length, array2.length);
        return joinedArray;
    }
    public static byte[] addAll(final byte[] array1, final byte... array2) {
        if (array1 == null) {
            return clone(array2);
        }
        if (array2 == null) {
            return clone(array1);
        }
        final byte[] joinedArray = new byte[addExact(array1.length, array2)];
        System.arraycopy(array1, 0, joinedArray, 0, array1.length);
        System.arraycopy(array2, 0, joinedArray, array1.length, array2.length);
        return joinedArray;
    }
    public static char[] addAll(final char[] array1, final char... array2) {
        if (array1 == null) {
            return clone(array2);
        }
        if (array2 == null) {
            return clone(array1);
        }
        final char[] joinedArray = new char[addExact(array1.length, array2)];
        System.arraycopy(array1, 0, joinedArray, 0, array1.length);
        System.arraycopy(array2, 0, joinedArray, array1.length, array2.length);
        return joinedArray;
    }
    public static double[] addAll(final double[] array1, final double... array2) {
        if (array1 == null) {
            return clone(array2);
        }
        if (array2 == null) {
            return clone(array1);
        }
        final double[] joinedArray = new double[addExact(array1.length, array2)];
        System.arraycopy(array1, 0, joinedArray, 0, array1.length);
        System.arraycopy(array2, 0, joinedArray, array1.length, array2.length);
        return joinedArray;
    }
    public static float[] addAll(final float[] array1, final float... array2) {
        if (array1 == null) {
            return clone(array2);
        }
        if (array2 == null) {
            return clone(array1);
        }
        final float[] joinedArray = new float[addExact(array1.length, array2)];
        System.arraycopy(array1, 0, joinedArray, 0, array1.length);
        System.arraycopy(array2, 0, joinedArray, array1.length, array2.length);
        return joinedArray;
    }
    public static int[] addAll(final int[] array1, final int... array2) {
        if (array1 == null) {
            return clone(array2);
        }
        if (array2 == null) {
            return clone(array1);
        }
        final int[] joinedArray = new int[addExact(array1.length, array2)];
        System.arraycopy(array1, 0, joinedArray, 0, array1.length);
        System.arraycopy(array2, 0, joinedArray, array1.length, array2.length);
        return joinedArray;
    }
    public static long[] addAll(final long[] array1, final long... array2) {
        if (array1 == null) {
            return clone(array2);
        }
        if (array2 == null) {
            return clone(array1);
        }
        final long[] joinedArray = new long[addExact(array1.length, array2)];
        System.arraycopy(array1, 0, joinedArray, 0, array1.length);
        System.arraycopy(array2, 0, joinedArray, array1.length, array2.length);
        return joinedArray;
    }
    public static short[] addAll(final short[] array1, final short... array2) {
        if (array1 == null) {
            return clone(array2);
        }
        if (array2 == null) {
            return clone(array1);
        }
        final short[] joinedArray = new short[addExact(array1.length, array2)];
        System.arraycopy(array1, 0, joinedArray, 0, array1.length);
        System.arraycopy(array2, 0, joinedArray, array1.length, array2.length);
        return joinedArray;
    }
    public static <T> T[] addAll(final T[] array1, @SuppressWarnings("unchecked") final T... array2) {
        if (array1 == null) {
            return clone(array2);
        }
        if (array2 == null) {
            return clone(array1);
        }
        final Class<T> type1 = getComponentType(array1);
        final T[] joinedArray = arraycopy(array1, 0, 0, array1.length, () -> newInstance(type1, addExact(array1.length, array2)));
        try {
            System.arraycopy(array2, 0, joinedArray, array1.length, array2.length);
        } catch (final ArrayStoreException ase) {
            final Class<?> type2 = array2.getClass().getComponentType();
            if (!type1.isAssignableFrom(type2)) {
                throw new IllegalArgumentException("Cannot store " + type2.getName() + " in an array of " + type1.getName(), ase);
            }
            throw ase; 
        }
        return joinedArray;
    }
    private static int addExact(final int totalLength, final Object array) {
        try {
            final int length = MathBridge.addExact(totalLength, getLength(array));
            if (length > SAFE_MAX_ARRAY_LENGTH) {
                throw new IllegalArgumentException("Total arrays length exceed " + SAFE_MAX_ARRAY_LENGTH);
            }
            return length;
        } catch (final ArithmeticException exception) {
            throw new IllegalArgumentException("Total arrays length exceed " + SAFE_MAX_ARRAY_LENGTH);
        }
    }
    public static boolean[] addFirst(final boolean[] array, final boolean element) {
        return array == null ? add(array, element) : insert(0, array, element);
    }
    public static byte[] addFirst(final byte[] array, final byte element) {
        return array == null ? add(array, element) : insert(0, array, element);
    }
    public static char[] addFirst(final char[] array, final char element) {
        return array == null ? add(array, element) : insert(0, array, element);
    }
    public static double[] addFirst(final double[] array, final double element) {
        return array == null ? add(array, element) : insert(0, array, element);
    }
    public static float[] addFirst(final float[] array, final float element) {
        return array == null ? add(array, element) : insert(0, array, element);
    }
    public static int[] addFirst(final int[] array, final int element) {
        return array == null ? add(array, element) : insert(0, array, element);
    }
    public static long[] addFirst(final long[] array, final long element) {
        return array == null ? add(array, element) : insert(0, array, element);
    }
    public static short[] addFirst(final short[] array, final short element) {
        return array == null ? add(array, element) : insert(0, array, element);
    }
    public static <T> T[] addFirst(final T[] array, final T element) {
        return array == null ? add(array, element) : insert(0, array, element);
    }
    public static <T> T arraycopy(final T source, final int sourcePos, final int destPos, final int length, final Function<Integer, T> allocator) {
        return arraycopy(source, sourcePos, allocator.apply(length), destPos, length);
    }
    public static <T> T arraycopy(final T source, final int sourcePos, final int destPos, final int length, final Supplier<T> allocator) {
        return arraycopy(source, sourcePos, allocator.get(), destPos, length);
    }
    public static <T> T arraycopy(final T source, final int sourcePos, final T dest, final int destPos, final int length) {
        System.arraycopy(source, sourcePos, dest, destPos, length);
        return dest;
    }
    public static boolean[] clone(final boolean[] array) {
        return array != null ? array.clone() : null;
    }
    public static byte[] clone(final byte[] array) {
        return array != null ? array.clone() : null;
    }
    public static char[] clone(final char[] array) {
        return array != null ? array.clone() : null;
    }
    public static double[] clone(final double[] array) {
        return array != null ? array.clone() : null;
    }
    public static float[] clone(final float[] array) {
        return array != null ? array.clone() : null;
    }
    public static int[] clone(final int[] array) {
        return array != null ? array.clone() : null;
    }
    public static long[] clone(final long[] array) {
        return array != null ? array.clone() : null;
    }
    public static short[] clone(final short[] array) {
        return array != null ? array.clone() : null;
    }
    public static <T> T[] clone(final T[] array) {
        return array != null ? array.clone() : null;
    }
    public static boolean[] concat(final boolean[]... arrays) {
        int totalLength = 0;
        for (final boolean[] array : arrays) {
            totalLength = addExact(totalLength, array);
        }
        final boolean[] result = new boolean[totalLength];
        int currentPos = 0;
        for (final boolean[] array : arrays) {
            if (array != null && array.length > 0) {
                System.arraycopy(array, 0, result, currentPos, array.length);
                currentPos += array.length;
            }
        }
        return result;
    }
    public static byte[] concat(final byte[]... arrays) {
        int totalLength = 0;
        for (final byte[] array : arrays) {
            totalLength = addExact(totalLength, array);
        }
        final byte[] result = new byte[totalLength];
        int currentPos = 0;
        for (final byte[] array : arrays) {
            if (array != null && array.length > 0) {
                System.arraycopy(array, 0, result, currentPos, array.length);
                currentPos += array.length;
            }
        }
        return result;
    }
    public static char[] concat(final char[]... arrays) {
        int totalLength = 0;
        for (final char[] array : arrays) {
            totalLength = addExact(totalLength, array);
        }
        final char[] result = new char[totalLength];
        int currentPos = 0;
        for (final char[] array : arrays) {
            if (array != null && array.length > 0) {
                System.arraycopy(array, 0, result, currentPos, array.length);
                currentPos += array.length;
            }
        }
        return result;
    }
    public static double[] concat(final double[]... arrays) {
        int totalLength = 0;
        for (final double[] array : arrays) {
            totalLength = addExact(totalLength, array);
        }
        final double[] result = new double[totalLength];
        int currentPos = 0;
        for (final double[] array : arrays) {
            if (array != null && array.length > 0) {
                System.arraycopy(array, 0, result, currentPos, array.length);
                currentPos += array.length;
            }
        }
        return result;
    }
    public static float[] concat(final float[]... arrays) {
        int totalLength = 0;
        for (final float[] array : arrays) {
            totalLength = addExact(totalLength, array);
        }
        final float[] result = new float[totalLength];
        int currentPos = 0;
        for (final float[] array : arrays) {
            if (array != null && array.length > 0) {
                System.arraycopy(array, 0, result, currentPos, array.length);
                currentPos += array.length;
            }
        }
        return result;
    }
    public static int[] concat(final int[]... arrays) {
        int totalLength = 0;
        for (final int[] array : arrays) {
            totalLength = addExact(totalLength, array);
        }
        final int[] result = new int[totalLength];
        int currentPos = 0;
        for (final int[] array : arrays) {
            if (array != null && array.length > 0) {
                System.arraycopy(array, 0, result, currentPos, array.length);
                currentPos += array.length;
            }
        }
        return result;
    }
    public static long[] concat(final long[]... arrays) {
        int totalLength = 0;
        for (final long[] array : arrays) {
            totalLength = addExact(totalLength, array);
        }
        final long[] result = new long[totalLength];
        int currentPos = 0;
        for (final long[] array : arrays) {
            if (array != null && array.length > 0) {
                System.arraycopy(array, 0, result, currentPos, array.length);
                currentPos += array.length;
            }
        }
        return result;
    }
    public static short[] concat(final short[]... arrays) {
        int totalLength = 0;
        for (final short[] array : arrays) {
            totalLength = addExact(totalLength, array);
        }
        final short[] result = new short[totalLength];
        int currentPos = 0;
        for (final short[] array : arrays) {
            if (array != null && array.length > 0) {
                System.arraycopy(array, 0, result, currentPos, array.length);
                currentPos += array.length;
            }
        }
        return result;
    }
    public static boolean contains(final boolean[] array, final boolean valueToFind) {
        return indexOf(array, valueToFind) != INDEX_NOT_FOUND;
    }
    public static boolean contains(final byte[] array, final byte valueToFind) {
        return indexOf(array, valueToFind) != INDEX_NOT_FOUND;
    }
    public static boolean contains(final char[] array, final char valueToFind) {
        return indexOf(array, valueToFind) != INDEX_NOT_FOUND;
    }
    public static boolean contains(final double[] array, final double valueToFind) {
        return indexOf(array, valueToFind) != INDEX_NOT_FOUND;
    }
    public static boolean contains(final double[] array, final double valueToFind, final double tolerance) {
        return indexOf(array, valueToFind, 0, tolerance) != INDEX_NOT_FOUND;
    }
    public static boolean contains(final float[] array, final float valueToFind) {
        return indexOf(array, valueToFind) != INDEX_NOT_FOUND;
    }
    public static boolean contains(final int[] array, final int valueToFind) {
        return indexOf(array, valueToFind) != INDEX_NOT_FOUND;
    }
    public static boolean contains(final long[] array, final long valueToFind) {
        return indexOf(array, valueToFind) != INDEX_NOT_FOUND;
    }
    public static boolean contains(final Object[] array, final Object objectToFind) {
        return indexOf(array, objectToFind) != INDEX_NOT_FOUND;
    }
    public static boolean contains(final short[] array, final short valueToFind) {
        return indexOf(array, valueToFind) != INDEX_NOT_FOUND;
    }
    public static boolean containsAny(final int[] array, final int... objectsToFind) {
        return IntStreams.of(objectsToFind).anyMatch(e -> contains(array, e));
    }
    public static boolean containsAny(final Object[] array, final Object... objectsToFind) {
        return Streams.of(objectsToFind).anyMatch(e -> contains(array, e));
    }
    private static Object copyArrayGrow1(final Object array, final Class<?> newArrayComponentType) {
        if (array != null) {
            final int arrayLength = Array.getLength(array);
            final Object newArray = Array.newInstance(array.getClass().getComponentType(), arrayLength + 1);
            System.arraycopy(array, 0, newArray, 0, arrayLength);
            return newArray;
        }
        return Array.newInstance(newArrayComponentType, 1);
    }
    public static <T> T get(final T[] array, final int index) {
        return get(array, index, null);
    }
    public static <T> T get(final T[] array, final int index, final T defaultValue) {
        return isArrayIndexValid(array, index) ? array[index] : defaultValue;
    }
    public static <T> Class<T> getComponentType(final T[] array) {
        return ClassUtils.getComponentType(ObjectUtils.getClass(array));
    }
    public static int getDimensions(final Object array) {
        int dimensions = 0;
        if (array != null) {
            Class<?> arrayClass = array.getClass();
            while (arrayClass.isArray()) {
                dimensions++;
                arrayClass = arrayClass.getComponentType();
            }
        }
        return dimensions;
    }
    public static int getLength(final Object array) {
        return array != null ? Array.getLength(array) : 0;
    }
    public static int hashCode(final Object array) {
        return new HashCodeBuilder().append(array).toHashCode();
    }
    static <K> void increment(final Map<K, MutableInt> occurrences, final K boxed) {
        occurrences.computeIfAbsent(boxed, k -> new MutableInt()).increment();
    }
    public static BitSet indexesOf(final boolean[] array, final boolean valueToFind) {
        return indexesOf(array, valueToFind, 0);
    }
    public static BitSet indexesOf(final boolean[] array, final boolean valueToFind, int startIndex) {
        final BitSet bitSet = new BitSet();
        if (array != null) {
            while (startIndex < array.length) {
                startIndex = indexOf(array, valueToFind, startIndex);
                if (startIndex == INDEX_NOT_FOUND) {
                    break;
                }
                bitSet.set(startIndex);
                ++startIndex;
            }
        }
        return bitSet;
    }
    public static BitSet indexesOf(final byte[] array, final byte valueToFind) {
        return indexesOf(array, valueToFind, 0);
    }
    public static BitSet indexesOf(final byte[] array, final byte valueToFind, int startIndex) {
        final BitSet bitSet = new BitSet();
        if (array != null) {
            while (startIndex < array.length) {
                startIndex = indexOf(array, valueToFind, startIndex);
                if (startIndex == INDEX_NOT_FOUND) {
                    break;
                }
                bitSet.set(startIndex);
                ++startIndex;
            }
        }
        return bitSet;
    }
    public static BitSet indexesOf(final char[] array, final char valueToFind) {
        return indexesOf(array, valueToFind, 0);
    }
    public static BitSet indexesOf(final char[] array, final char valueToFind, int startIndex) {
        final BitSet bitSet = new BitSet();
        if (array != null) {
            while (startIndex < array.length) {
                startIndex = indexOf(array, valueToFind, startIndex);
                if (startIndex == INDEX_NOT_FOUND) {
                    break;
                }
                bitSet.set(startIndex);
                ++startIndex;
            }
        }
        return bitSet;
    }
    public static BitSet indexesOf(final double[] array, final double valueToFind) {
        return indexesOf(array, valueToFind, 0);
    }
    public static BitSet indexesOf(final double[] array, final double valueToFind, final double tolerance) {
        return indexesOf(array, valueToFind, 0, tolerance);
    }
    public static BitSet indexesOf(final double[] array, final double valueToFind, int startIndex) {
        final BitSet bitSet = new BitSet();
        if (array != null) {
            while (startIndex < array.length) {
                startIndex = indexOf(array, valueToFind, startIndex);
                if (startIndex == INDEX_NOT_FOUND) {
                    break;
                }
                bitSet.set(startIndex);
                ++startIndex;
            }
        }
        return bitSet;
    }
    public static BitSet indexesOf(final double[] array, final double valueToFind, int startIndex, final double tolerance) {
        final BitSet bitSet = new BitSet();
        if (array != null) {
            while (startIndex < array.length) {
                startIndex = indexOf(array, valueToFind, startIndex, tolerance);
                if (startIndex == INDEX_NOT_FOUND) {
                    break;
                }
                bitSet.set(startIndex);
                ++startIndex;
            }
        }
        return bitSet;
    }
    public static BitSet indexesOf(final float[] array, final float valueToFind) {
        return indexesOf(array, valueToFind, 0);
    }
    public static BitSet indexesOf(final float[] array, final float valueToFind, int startIndex) {
        final BitSet bitSet = new BitSet();
        if (array != null) {
            while (startIndex < array.length) {
                startIndex = indexOf(array, valueToFind, startIndex);
                if (startIndex == INDEX_NOT_FOUND) {
                    break;
                }
                bitSet.set(startIndex);
                ++startIndex;
            }
        }
        return bitSet;
    }
    public static BitSet indexesOf(final int[] array, final int valueToFind) {
        return indexesOf(array, valueToFind, 0);
    }
    public static BitSet indexesOf(final int[] array, final int valueToFind, int startIndex) {
        final BitSet bitSet = new BitSet();
        if (array != null) {
            while (startIndex < array.length) {
                startIndex = indexOf(array, valueToFind, startIndex);
                if (startIndex == INDEX_NOT_FOUND) {
                    break;
                }
                bitSet.set(startIndex);
                ++startIndex;
            }
        }
        return bitSet;
    }
    public static BitSet indexesOf(final long[] array, final long valueToFind) {
        return indexesOf(array, valueToFind, 0);
    }
    public static BitSet indexesOf(final long[] array, final long valueToFind, int startIndex) {
        final BitSet bitSet = new BitSet();
        if (array != null) {
            while (startIndex < array.length) {
                startIndex = indexOf(array, valueToFind, startIndex);
                if (startIndex == INDEX_NOT_FOUND) {
                    break;
                }
                bitSet.set(startIndex);
                ++startIndex;
            }
        }
        return bitSet;
    }
    public static BitSet indexesOf(final Object[] array, final Object objectToFind) {
        return indexesOf(array, objectToFind, 0);
    }
    public static BitSet indexesOf(final Object[] array, final Object objectToFind, int startIndex) {
        final BitSet bitSet = new BitSet();
        if (array != null) {
            while (startIndex < array.length) {
                startIndex = indexOf(array, objectToFind, startIndex);
                if (startIndex == INDEX_NOT_FOUND) {
                    break;
                }
                bitSet.set(startIndex);
                ++startIndex;
            }
        }
        return bitSet;
    }
    public static BitSet indexesOf(final short[] array, final short valueToFind) {
        return indexesOf(array, valueToFind, 0);
    }
    public static BitSet indexesOf(final short[] array, final short valueToFind, int startIndex) {
        final BitSet bitSet = new BitSet();
        if (array != null) {
            while (startIndex < array.length) {
                startIndex = indexOf(array, valueToFind, startIndex);
                if (startIndex == INDEX_NOT_FOUND) {
                    break;
                }
                bitSet.set(startIndex);
                ++startIndex;
            }
        }
        return bitSet;
    }
    public static int indexOf(final boolean[] array, final boolean valueToFind) {
        return indexOf(array, valueToFind, 0);
    }
    public static int indexOf(final boolean[] array, final boolean valueToFind, final int startIndex) {
        if (isEmpty(array)) {
            return INDEX_NOT_FOUND;
        }
        for (int i = max0(startIndex); i < array.length; i++) {
            if (valueToFind == array[i]) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int indexOf(final byte[] array, final byte valueToFind) {
        return indexOf(array, valueToFind, 0);
    }
    public static int indexOf(final byte[] array, final byte valueToFind, final int startIndex) {
        if (isEmpty(array)) {
            return INDEX_NOT_FOUND;
        }
        for (int i = max0(startIndex); i < array.length; i++) {
            if (valueToFind == array[i]) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int indexOf(final char[] array, final char valueToFind) {
        return indexOf(array, valueToFind, 0);
    }
    public static int indexOf(final char[] array, final char valueToFind, final int startIndex) {
        if (isEmpty(array)) {
            return INDEX_NOT_FOUND;
        }
        for (int i = max0(startIndex); i < array.length; i++) {
            if (valueToFind == array[i]) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int indexOf(final double[] array, final double valueToFind) {
        return indexOf(array, valueToFind, 0);
    }
    public static int indexOf(final double[] array, final double valueToFind, final double tolerance) {
        return indexOf(array, valueToFind, 0, tolerance);
    }
    public static int indexOf(final double[] array, final double valueToFind, final int startIndex) {
        if (Double.isNaN(valueToFind)) {
            return indexOfNaN(array, startIndex);
        }
        if (isEmpty(array)) {
            return INDEX_NOT_FOUND;
        }
        for (int i = max0(startIndex); i < array.length; i++) {
            if (valueToFind == array[i]) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int indexOf(final double[] array, final double valueToFind, final int startIndex, final double tolerance) {
        if (Double.isNaN(valueToFind)) {
            return indexOfNaN(array, startIndex);
        }
        if (isEmpty(array)) {
            return INDEX_NOT_FOUND;
        }
        final double min = valueToFind - tolerance;
        final double max = valueToFind + tolerance;
        for (int i = max0(startIndex); i < array.length; i++) {
            if (array[i] >= min && array[i] <= max) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int indexOf(final float[] array, final float valueToFind) {
        return indexOf(array, valueToFind, 0);
    }
    public static int indexOf(final float[] array, final float valueToFind, final int startIndex) {
        if (isEmpty(array)) {
            return INDEX_NOT_FOUND;
        }
        final boolean searchNaN = Float.isNaN(valueToFind);
        for (int i = max0(startIndex); i < array.length; i++) {
            final float element = array[i];
            if (valueToFind == element || searchNaN && Float.isNaN(element)) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int indexOf(final int[] array, final int valueToFind) {
        return indexOf(array, valueToFind, 0);
    }
    public static int indexOf(final int[] array, final int valueToFind, final int startIndex) {
        if (isEmpty(array)) {
            return INDEX_NOT_FOUND;
        }
        for (int i = max0(startIndex); i < array.length; i++) {
            if (valueToFind == array[i]) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int indexOf(final long[] array, final long valueToFind) {
        return indexOf(array, valueToFind, 0);
    }
    public static int indexOf(final long[] array, final long valueToFind, final int startIndex) {
        if (isEmpty(array)) {
            return INDEX_NOT_FOUND;
        }
        for (int i = max0(startIndex); i < array.length; i++) {
            if (valueToFind == array[i]) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int indexOf(final Object[] array, final Object objectToFind) {
        return indexOf(array, objectToFind, 0);
    }
    public static int indexOf(final Object[] array, final Object objectToFind, int startIndex) {
        if (isEmpty(array)) {
            return INDEX_NOT_FOUND;
        }
        startIndex = max0(startIndex);
        if (objectToFind == null) {
            for (int i = startIndex; i < array.length; i++) {
                if (array[i] == null) {
                    return i;
                }
            }
        } else {
            for (int i = startIndex; i < array.length; i++) {
                if (objectToFind.equals(array[i])) {
                    return i;
                }
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int indexOf(final short[] array, final short valueToFind) {
        return indexOf(array, valueToFind, 0);
    }
    public static int indexOf(final short[] array, final short valueToFind, final int startIndex) {
        if (isEmpty(array)) {
            return INDEX_NOT_FOUND;
        }
        for (int i = max0(startIndex); i < array.length; i++) {
            if (valueToFind == array[i]) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    private static int indexOfNaN(final double[] array, final int startIndex) {
        if (isEmpty(array)) {
            return INDEX_NOT_FOUND;
        }
        for (int i = max0(startIndex); i < array.length; i++) {
            if (Double.isNaN(array[i])) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static boolean[] insert(final int index, final boolean[] array, final boolean... values) {
        if (array == null) {
            return null;
        }
        if (isEmpty(values)) {
            return clone(array);
        }
        if (index < 0 || index > array.length) {
            throw new IndexOutOfBoundsException("Index: " + index + ", Length: " + array.length);
        }
        final boolean[] result = new boolean[addExact(array.length, values)];
        System.arraycopy(values, 0, result, index, values.length);
        if (index > 0) {
            System.arraycopy(array, 0, result, 0, index);
        }
        if (index < array.length) {
            System.arraycopy(array, index, result, index + values.length, array.length - index);
        }
        return result;
    }
    public static byte[] insert(final int index, final byte[] array, final byte... values) {
        if (array == null) {
            return null;
        }
        if (isEmpty(values)) {
            return clone(array);
        }
        if (index < 0 || index > array.length) {
            throw new IndexOutOfBoundsException("Index: " + index + ", Length: " + array.length);
        }
        final byte[] result = new byte[addExact(array.length, values)];
        System.arraycopy(values, 0, result, index, values.length);
        if (index > 0) {
            System.arraycopy(array, 0, result, 0, index);
        }
        if (index < array.length) {
            System.arraycopy(array, index, result, index + values.length, array.length - index);
        }
        return result;
    }
    public static char[] insert(final int index, final char[] array, final char... values) {
        if (array == null) {
            return null;
        }
        if (isEmpty(values)) {
            return clone(array);
        }
        if (index < 0 || index > array.length) {
            throw new IndexOutOfBoundsException("Index: " + index + ", Length: " + array.length);
        }
        final char[] result = new char[addExact(array.length, values)];
        System.arraycopy(values, 0, result, index, values.length);
        if (index > 0) {
            System.arraycopy(array, 0, result, 0, index);
        }
        if (index < array.length) {
            System.arraycopy(array, index, result, index + values.length, array.length - index);
        }
        return result;
    }
    public static double[] insert(final int index, final double[] array, final double... values) {
        if (array == null) {
            return null;
        }
        if (isEmpty(values)) {
            return clone(array);
        }
        if (index < 0 || index > array.length) {
            throw new IndexOutOfBoundsException("Index: " + index + ", Length: " + array.length);
        }
        final double[] result = new double[addExact(array.length, values)];
        System.arraycopy(values, 0, result, index, values.length);
        if (index > 0) {
            System.arraycopy(array, 0, result, 0, index);
        }
        if (index < array.length) {
            System.arraycopy(array, index, result, index + values.length, array.length - index);
        }
        return result;
    }
    public static float[] insert(final int index, final float[] array, final float... values) {
        if (array == null) {
            return null;
        }
        if (isEmpty(values)) {
            return clone(array);
        }
        if (index < 0 || index > array.length) {
            throw new IndexOutOfBoundsException("Index: " + index + ", Length: " + array.length);
        }
        final float[] result = new float[addExact(array.length, values)];
        System.arraycopy(values, 0, result, index, values.length);
        if (index > 0) {
            System.arraycopy(array, 0, result, 0, index);
        }
        if (index < array.length) {
            System.arraycopy(array, index, result, index + values.length, array.length - index);
        }
        return result;
    }
    public static int[] insert(final int index, final int[] array, final int... values) {
        if (array == null) {
            return null;
        }
        if (isEmpty(values)) {
            return clone(array);
        }
        if (index < 0 || index > array.length) {
            throw new IndexOutOfBoundsException("Index: " + index + ", Length: " + array.length);
        }
        final int[] result = new int[addExact(array.length, values)];
        System.arraycopy(values, 0, result, index, values.length);
        if (index > 0) {
            System.arraycopy(array, 0, result, 0, index);
        }
        if (index < array.length) {
            System.arraycopy(array, index, result, index + values.length, array.length - index);
        }
        return result;
    }
    public static long[] insert(final int index, final long[] array, final long... values) {
        if (array == null) {
            return null;
        }
        if (isEmpty(values)) {
            return clone(array);
        }
        if (index < 0 || index > array.length) {
            throw new IndexOutOfBoundsException("Index: " + index + ", Length: " + array.length);
        }
        final long[] result = new long[addExact(array.length, values)];
        System.arraycopy(values, 0, result, index, values.length);
        if (index > 0) {
            System.arraycopy(array, 0, result, 0, index);
        }
        if (index < array.length) {
            System.arraycopy(array, index, result, index + values.length, array.length - index);
        }
        return result;
    }
    public static short[] insert(final int index, final short[] array, final short... values) {
        if (array == null) {
            return null;
        }
        if (isEmpty(values)) {
            return clone(array);
        }
        if (index < 0 || index > array.length) {
            throw new IndexOutOfBoundsException("Index: " + index + ", Length: " + array.length);
        }
        final short[] result = new short[addExact(array.length, values)];
        System.arraycopy(values, 0, result, index, values.length);
        if (index > 0) {
            System.arraycopy(array, 0, result, 0, index);
        }
        if (index < array.length) {
            System.arraycopy(array, index, result, index + values.length, array.length - index);
        }
        return result;
    }
    @SafeVarargs
    public static <T> T[] insert(final int index, final T[] array, final T... values) {
        if (array == null) {
            return null;
        }
        if (isEmpty(values)) {
            return clone(array);
        }
        if (index < 0 || index > array.length) {
            throw new IndexOutOfBoundsException("Index: " + index + ", Length: " + array.length);
        }
        final Class<T> type = getComponentType(array);
        final int length = addExact(array.length, values);
        final T[] result = newInstance(type, length);
        System.arraycopy(values, 0, result, index, values.length);
        if (index > 0) {
            System.arraycopy(array, 0, result, 0, index);
        }
        if (index < array.length) {
            System.arraycopy(array, index, result, index + values.length, array.length - index);
        }
        return result;
    }
    private static boolean isArrayEmpty(final Object array) {
        return getLength(array) == 0;
    }
    public static <T> boolean isArrayIndexValid(final T[] array, final int index) {
        return index >= 0 && getLength(array) > index;
    }
    public static boolean isEmpty(final boolean[] array) {
        return isArrayEmpty(array);
    }
    public static boolean isEmpty(final byte[] array) {
        return isArrayEmpty(array);
    }
    public static boolean isEmpty(final char[] array) {
        return isArrayEmpty(array);
    }
    public static boolean isEmpty(final double[] array) {
        return isArrayEmpty(array);
    }
    public static boolean isEmpty(final float[] array) {
        return isArrayEmpty(array);
    }
    public static boolean isEmpty(final int[] array) {
        return isArrayEmpty(array);
    }
    public static boolean isEmpty(final long[] array) {
        return isArrayEmpty(array);
    }
    public static boolean isEmpty(final Object[] array) {
        return isArrayEmpty(array);
    }
    public static boolean isEmpty(final short[] array) {
        return isArrayEmpty(array);
    }
    @Deprecated
    public static boolean isEquals(final Object array1, final Object array2) {
        return new EqualsBuilder().append(array1, array2).isEquals();
    }
    public static boolean isNotEmpty(final boolean[] array) {
        return !isEmpty(array);
    }
    public static boolean isNotEmpty(final byte[] array) {
        return !isEmpty(array);
    }
    public static boolean isNotEmpty(final char[] array) {
        return !isEmpty(array);
    }
    public static boolean isNotEmpty(final double[] array) {
        return !isEmpty(array);
    }
    public static boolean isNotEmpty(final float[] array) {
        return !isEmpty(array);
    }
    public static boolean isNotEmpty(final int[] array) {
        return !isEmpty(array);
    }
    public static boolean isNotEmpty(final long[] array) {
        return !isEmpty(array);
    }
    public static boolean isNotEmpty(final short[] array) {
        return !isEmpty(array);
    }
     public static <T> boolean isNotEmpty(final T[] array) {
         return !isEmpty(array);
     }
     public static boolean isSameLength(final boolean[] array1, final boolean[] array2) {
        return getLength(array1) == getLength(array2);
    }
    public static boolean isSameLength(final byte[] array1, final byte[] array2) {
        return getLength(array1) == getLength(array2);
    }
    public static boolean isSameLength(final char[] array1, final char[] array2) {
        return getLength(array1) == getLength(array2);
    }
    public static boolean isSameLength(final double[] array1, final double[] array2) {
        return getLength(array1) == getLength(array2);
    }
    public static boolean isSameLength(final float[] array1, final float[] array2) {
        return getLength(array1) == getLength(array2);
    }
    public static boolean isSameLength(final int[] array1, final int[] array2) {
        return getLength(array1) == getLength(array2);
    }
    public static boolean isSameLength(final long[] array1, final long[] array2) {
        return getLength(array1) == getLength(array2);
    }
    public static boolean isSameLength(final Object array1, final Object array2) {
        return getLength(array1) == getLength(array2);
    }
    public static boolean isSameLength(final Object[] array1, final Object[] array2) {
        return getLength(array1) == getLength(array2);
    }
    public static boolean isSameLength(final short[] array1, final short[] array2) {
        return getLength(array1) == getLength(array2);
    }
    public static boolean isSameType(final Object array1, final Object array2) {
        if (array1 == null || array2 == null) {
            throw new IllegalArgumentException("The Array must not be null");
        }
        return array1.getClass().getName().equals(array2.getClass().getName());
    }
    public static boolean isSorted(final boolean[] array) {
        if (getLength(array) < 2) {
            return true;
        }
        boolean previous = array[0];
        final int n = array.length;
        for (int i = 1; i < n; i++) {
            final boolean current = array[i];
            if (BooleanUtils.compare(previous, current) > 0) {
                return false;
            }
            previous = current;
        }
        return true;
    }
    public static boolean isSorted(final byte[] array) {
        if (getLength(array) < 2) {
            return true;
        }
        byte previous = array[0];
        final int n = array.length;
        for (int i = 1; i < n; i++) {
            final byte current = array[i];
            if (Byte.compare(previous, current) > 0) {
                return false;
            }
            previous = current;
        }
        return true;
    }
    public static boolean isSorted(final char[] array) {
        if (getLength(array) < 2) {
            return true;
        }
        char previous = array[0];
        final int n = array.length;
        for (int i = 1; i < n; i++) {
            final char current = array[i];
            if (CharUtils.compare(previous, current) > 0) {
                return false;
            }
            previous = current;
        }
        return true;
    }
    public static boolean isSorted(final double[] array) {
        if (getLength(array) < 2) {
            return true;
        }
        double previous = array[0];
        final int n = array.length;
        for (int i = 1; i < n; i++) {
            final double current = array[i];
            if (Double.compare(previous, current) > 0) {
                return false;
            }
            previous = current;
        }
        return true;
    }
    public static boolean isSorted(final float[] array) {
        if (getLength(array) < 2) {
            return true;
        }
        float previous = array[0];
        final int n = array.length;
        for (int i = 1; i < n; i++) {
            final float current = array[i];
            if (Float.compare(previous, current) > 0) {
                return false;
            }
            previous = current;
        }
        return true;
    }
    public static boolean isSorted(final int[] array) {
        if (getLength(array) < 2) {
            return true;
        }
        int previous = array[0];
        final int n = array.length;
        for (int i = 1; i < n; i++) {
            final int current = array[i];
            if (Integer.compare(previous, current) > 0) {
                return false;
            }
            previous = current;
        }
        return true;
    }
    public static boolean isSorted(final long[] array) {
        if (getLength(array) < 2) {
            return true;
        }
        long previous = array[0];
        final int n = array.length;
        for (int i = 1; i < n; i++) {
            final long current = array[i];
            if (Long.compare(previous, current) > 0) {
                return false;
            }
            previous = current;
        }
        return true;
    }
    public static boolean isSorted(final short[] array) {
        if (getLength(array) < 2) {
            return true;
        }
        short previous = array[0];
        final int n = array.length;
        for (int i = 1; i < n; i++) {
            final short current = array[i];
            if (Short.compare(previous, current) > 0) {
                return false;
            }
            previous = current;
        }
        return true;
    }
    public static <T extends Comparable<? super T>> boolean isSorted(final T[] array) {
        return isSorted(array, Comparable::compareTo);
    }
    public static <T> boolean isSorted(final T[] array, final Comparator<T> comparator) {
        Objects.requireNonNull(comparator, "comparator");
        if (getLength(array) < 2) {
            return true;
        }
        T previous = array[0];
        final int n = array.length;
        for (int i = 1; i < n; i++) {
            final T current = array[i];
            if (comparator.compare(previous, current) > 0) {
                return false;
            }
            previous = current;
        }
        return true;
    }
    public static int lastIndexOf(final boolean[] array, final boolean valueToFind) {
        return lastIndexOf(array, valueToFind, Integer.MAX_VALUE);
    }
    public static int lastIndexOf(final boolean[] array, final boolean valueToFind, int startIndex) {
        if (isEmpty(array) || startIndex < 0) {
            return INDEX_NOT_FOUND;
        }
        if (startIndex >= array.length) {
            startIndex = array.length - 1;
        }
        for (int i = startIndex; i >= 0; i--) {
            if (valueToFind == array[i]) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int lastIndexOf(final byte[] array, final byte valueToFind) {
        return lastIndexOf(array, valueToFind, Integer.MAX_VALUE);
    }
    public static int lastIndexOf(final byte[] array, final byte valueToFind, int startIndex) {
        if (array == null || startIndex < 0) {
            return INDEX_NOT_FOUND;
        }
        if (startIndex >= array.length) {
            startIndex = array.length - 1;
        }
        for (int i = startIndex; i >= 0; i--) {
            if (valueToFind == array[i]) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int lastIndexOf(final char[] array, final char valueToFind) {
        return lastIndexOf(array, valueToFind, Integer.MAX_VALUE);
    }
    public static int lastIndexOf(final char[] array, final char valueToFind, int startIndex) {
        if (array == null || startIndex < 0) {
            return INDEX_NOT_FOUND;
        }
        if (startIndex >= array.length) {
            startIndex = array.length - 1;
        }
        for (int i = startIndex; i >= 0; i--) {
            if (valueToFind == array[i]) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int lastIndexOf(final double[] array, final double valueToFind) {
        return lastIndexOf(array, valueToFind, Integer.MAX_VALUE);
    }
    public static int lastIndexOf(final double[] array, final double valueToFind, final double tolerance) {
        return lastIndexOf(array, valueToFind, Integer.MAX_VALUE, tolerance);
    }
    public static int lastIndexOf(final double[] array, final double valueToFind, int startIndex) {
        if (Double.isNaN(valueToFind)) {
            return lastIndexOfNaN(array, startIndex);
        }
        if (isEmpty(array) || startIndex < 0) {
            return INDEX_NOT_FOUND;
        }
        if (startIndex >= array.length) {
            startIndex = array.length - 1;
        }
        for (int i = startIndex; i >= 0; i--) {
            if (valueToFind == array[i]) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int lastIndexOf(final double[] array, final double valueToFind, int startIndex, final double tolerance) {
        if (Double.isNaN(valueToFind)) {
            return lastIndexOfNaN(array, startIndex);
        }
        if (isEmpty(array) || startIndex < 0) {
            return INDEX_NOT_FOUND;
        }
        if (startIndex >= array.length) {
            startIndex = array.length - 1;
        }
        final double min = valueToFind - tolerance;
        final double max = valueToFind + tolerance;
        for (int i = startIndex; i >= 0; i--) {
            if (array[i] >= min && array[i] <= max) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int lastIndexOf(final float[] array, final float valueToFind) {
        return lastIndexOf(array, valueToFind, Integer.MAX_VALUE);
    }
    public static int lastIndexOf(final float[] array, final float valueToFind, int startIndex) {
        if (isEmpty(array) || startIndex < 0) {
            return INDEX_NOT_FOUND;
        }
        if (startIndex >= array.length) {
            startIndex = array.length - 1;
        }
        final boolean searchNaN = Float.isNaN(valueToFind);
        for (int i = startIndex; i >= 0; i--) {
            final float element = array[i];
            if (valueToFind == element || searchNaN && Float.isNaN(element)) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int lastIndexOf(final int[] array, final int valueToFind) {
        return lastIndexOf(array, valueToFind, Integer.MAX_VALUE);
    }
    public static int lastIndexOf(final int[] array, final int valueToFind, int startIndex) {
        if (array == null || startIndex < 0) {
            return INDEX_NOT_FOUND;
        }
        if (startIndex >= array.length) {
            startIndex = array.length - 1;
        }
        for (int i = startIndex; i >= 0; i--) {
            if (valueToFind == array[i]) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int lastIndexOf(final long[] array, final long valueToFind) {
        return lastIndexOf(array, valueToFind, Integer.MAX_VALUE);
    }
    public static int lastIndexOf(final long[] array, final long valueToFind, int startIndex) {
        if (array == null || startIndex < 0) {
            return INDEX_NOT_FOUND;
        }
        if (startIndex >= array.length) {
            startIndex = array.length - 1;
        }
        for (int i = startIndex; i >= 0; i--) {
            if (valueToFind == array[i]) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int lastIndexOf(final Object[] array, final Object objectToFind) {
        return lastIndexOf(array, objectToFind, Integer.MAX_VALUE);
    }
    public static int lastIndexOf(final Object[] array, final Object objectToFind, int startIndex) {
        if (array == null || startIndex < 0) {
            return INDEX_NOT_FOUND;
        }
        if (startIndex >= array.length) {
            startIndex = array.length - 1;
        }
        if (objectToFind == null) {
            for (int i = startIndex; i >= 0; i--) {
                if (array[i] == null) {
                    return i;
                }
            }
        } else if (array.getClass().getComponentType().isInstance(objectToFind)) {
            for (int i = startIndex; i >= 0; i--) {
                if (objectToFind.equals(array[i])) {
                    return i;
                }
            }
        }
        return INDEX_NOT_FOUND;
    }
    public static int lastIndexOf(final short[] array, final short valueToFind) {
        return lastIndexOf(array, valueToFind, Integer.MAX_VALUE);
    }
    public static int lastIndexOf(final short[] array, final short valueToFind, int startIndex) {
        if (array == null || startIndex < 0) {
            return INDEX_NOT_FOUND;
        }
        if (startIndex >= array.length) {
            startIndex = array.length - 1;
        }
        for (int i = startIndex; i >= 0; i--) {
            if (valueToFind == array[i]) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    private static int lastIndexOfNaN(final double[] array, final int startIndex) {
        if (isEmpty(array) || startIndex < 0) {
            return INDEX_NOT_FOUND;
        }
        for (int i = Math.min(startIndex, array.length - 1); i >= 0; i--) {
            if (Double.isNaN(array[i])) {
                return i;
            }
        }
        return INDEX_NOT_FOUND;
    }
    private static <T, R, E extends Throwable> R[] map(final T[] array, final Class<R> componentType, final FailableFunction<? super T, ? extends R, E> mapper)
            throws E {
        return ArrayFill.fill(newInstance(componentType, array.length), i -> mapper.apply(array[i]));
    }
    private static int max0(final int other) {
        return Math.max(0, other);
    }
    @SuppressWarnings("unchecked") 
    public static <T> T[] newInstance(final Class<T> componentType, final int length) {
        return (T[]) Array.newInstance(componentType, length);
    }
    public static <T> T[] nullTo(final T[] array, final T[] defaultArray) {
        return isEmpty(array) ? defaultArray : array;
    }
    public static boolean[] nullToEmpty(final boolean[] array) {
        return isEmpty(array) ? EMPTY_BOOLEAN_ARRAY : array;
    }
    public static Boolean[] nullToEmpty(final Boolean[] array) {
        return nullTo(array, EMPTY_BOOLEAN_OBJECT_ARRAY);
    }
    public static byte[] nullToEmpty(final byte[] array) {
        return isEmpty(array) ? EMPTY_BYTE_ARRAY : array;
    }
    public static Byte[] nullToEmpty(final Byte[] array) {
        return nullTo(array, EMPTY_BYTE_OBJECT_ARRAY);
    }
    public static char[] nullToEmpty(final char[] array) {
        return isEmpty(array) ? EMPTY_CHAR_ARRAY : array;
    }
    public static Character[] nullToEmpty(final Character[] array) {
        return nullTo(array, EMPTY_CHARACTER_OBJECT_ARRAY);
    }
    public static Class<?>[] nullToEmpty(final Class<?>[] array) {
        return nullTo(array, EMPTY_CLASS_ARRAY);
    }
    public static double[] nullToEmpty(final double[] array) {
        return isEmpty(array) ? EMPTY_DOUBLE_ARRAY : array;
    }
    public static Double[] nullToEmpty(final Double[] array) {
        return nullTo(array, EMPTY_DOUBLE_OBJECT_ARRAY);
    }
    public static float[] nullToEmpty(final float[] array) {
        return isEmpty(array) ? EMPTY_FLOAT_ARRAY : array;
    }
    public static Float[] nullToEmpty(final Float[] array) {
        return nullTo(array, EMPTY_FLOAT_OBJECT_ARRAY);
    }
    public static int[] nullToEmpty(final int[] array) {
        return isEmpty(array) ? EMPTY_INT_ARRAY : array;
    }
    public static Integer[] nullToEmpty(final Integer[] array) {
        return nullTo(array, EMPTY_INTEGER_OBJECT_ARRAY);
    }
    public static long[] nullToEmpty(final long[] array) {
        return isEmpty(array) ? EMPTY_LONG_ARRAY : array;
    }
    public static Long[] nullToEmpty(final Long[] array) {
        return nullTo(array, EMPTY_LONG_OBJECT_ARRAY);
    }
    public static Object[] nullToEmpty(final Object[] array) {
        return nullTo(array, EMPTY_OBJECT_ARRAY);
    }
    public static short[] nullToEmpty(final short[] array) {
        return isEmpty(array) ? EMPTY_SHORT_ARRAY : array;
    }
    public static Short[] nullToEmpty(final Short[] array) {
        return nullTo(array, EMPTY_SHORT_OBJECT_ARRAY);
    }
    public static String[] nullToEmpty(final String[] array) {
        return nullTo(array, EMPTY_STRING_ARRAY);
    }
    public static <T> T[] nullToEmpty(final T[] array, final Class<T[]> type) {
        if (type == null) {
            throw new IllegalArgumentException("The type must not be null");
        }
        if (array == null) {
            return type.cast(Array.newInstance(type.getComponentType(), 0));
        }
        return array;
    }
    private static ThreadLocalRandom random() {
        return ThreadLocalRandom.current();
    }
    public static boolean[] remove(final boolean[] array, final int index) {
        return (boolean[]) remove((Object) array, index);
    }
    public static byte[] remove(final byte[] array, final int index) {
        return (byte[]) remove((Object) array, index);
    }
    public static char[] remove(final char[] array, final int index) {
        return (char[]) remove((Object) array, index);
    }
    public static double[] remove(final double[] array, final int index) {
        return (double[]) remove((Object) array, index);
    }
    public static float[] remove(final float[] array, final int index) {
        return (float[]) remove((Object) array, index);
    }
    public static int[] remove(final int[] array, final int index) {
        return (int[]) remove((Object) array, index);
    }
    public static long[] remove(final long[] array, final int index) {
        return (long[]) remove((Object) array, index);
    }
    private static Object remove(final Object array, final int index) {
        final int length = getLength(array);
        if (index < 0 || index >= length) {
            throw new IndexOutOfBoundsException("Index: " + index + ", Length: " + length);
        }
        final Object result = Array.newInstance(array.getClass().getComponentType(), length - 1);
        System.arraycopy(array, 0, result, 0, index);
        if (index < length - 1) {
            System.arraycopy(array, index + 1, result, index, length - index - 1);
        }
        return result;
    }
    public static short[] remove(final short[] array, final int index) {
        return (short[]) remove((Object) array, index);
    }
    @SuppressWarnings("unchecked") 
    public static <T> T[] remove(final T[] array, final int index) {
        return (T[]) remove((Object) array, index);
    }
    public static boolean[] removeAll(final boolean[] array, final int... indices) {
        return (boolean[]) removeAll((Object) array, indices);
    }
    public static byte[] removeAll(final byte[] array, final int... indices) {
        return (byte[]) removeAll((Object) array, indices);
    }
    public static char[] removeAll(final char[] array, final int... indices) {
        return (char[]) removeAll((Object) array, indices);
    }
    public static double[] removeAll(final double[] array, final int... indices) {
        return (double[]) removeAll((Object) array, indices);
    }
    public static float[] removeAll(final float[] array, final int... indices) {
        return (float[]) removeAll((Object) array, indices);
    }
    public static int[] removeAll(final int[] array, final int... indices) {
        return (int[]) removeAll((Object) array, indices);
    }
    public static long[] removeAll(final long[] array, final int... indices) {
        return (long[]) removeAll((Object) array, indices);
    }
    static Object removeAll(final Object array, final int... indices) {
        if (array == null) {
            return null;
        }
        final int length = getLength(array);
        int diff = 0; 
        final int[] clonedIndices = ArraySorter.sort(clone(indices));
        if (isNotEmpty(clonedIndices)) {
            int i = clonedIndices.length;
            int prevIndex = length;
            while (--i >= 0) {
                final int index = clonedIndices[i];
                if (index < 0 || index >= length) {
                    throw new IndexOutOfBoundsException("Index: " + index + ", Length: " + length);
                }
                if (index >= prevIndex) {
                    continue;
                }
                diff++;
                prevIndex = index;
            }
        }
        final Object result = Array.newInstance(array.getClass().getComponentType(), length - diff);
        if (diff < length && clonedIndices != null) {
            int end = length; 
            int dest = length - diff; 
            for (int i = clonedIndices.length - 1; i >= 0; i--) {
                final int index = clonedIndices[i];
                if (end - index > 1) { 
                    final int cp = end - index - 1;
                    dest -= cp;
                    System.arraycopy(array, index + 1, result, dest, cp);
                }
                end = index;
            }
            if (end > 0) {
                System.arraycopy(array, 0, result, 0, end);
            }
        }
        return result;
    }
    public static short[] removeAll(final short[] array, final int... indices) {
        return (short[]) removeAll((Object) array, indices);
    }
    @SuppressWarnings("unchecked") 
    public static <T> T[] removeAll(final T[] array, final int... indices) {
        return (T[]) removeAll((Object) array, indices);
    }
    @Deprecated
    public static boolean[] removeAllOccurences(final boolean[] array, final boolean element) {
        return (boolean[]) removeAt(array, indexesOf(array, element));
    }
    @Deprecated
    public static byte[] removeAllOccurences(final byte[] array, final byte element) {
        return (byte[]) removeAt(array, indexesOf(array, element));
    }
    @Deprecated
    public static char[] removeAllOccurences(final char[] array, final char element) {
        return (char[]) removeAt(array, indexesOf(array, element));
    }
    @Deprecated
    public static double[] removeAllOccurences(final double[] array, final double element) {
        return (double[]) removeAt(array, indexesOf(array, element));
    }
    @Deprecated
    public static float[] removeAllOccurences(final float[] array, final float element) {
        return (float[]) removeAt(array, indexesOf(array, element));
    }
    @Deprecated
    public static int[] removeAllOccurences(final int[] array, final int element) {
        return (int[]) removeAt(array, indexesOf(array, element));
    }
    @Deprecated
    public static long[] removeAllOccurences(final long[] array, final long element) {
        return (long[]) removeAt(array, indexesOf(array, element));
    }
    @Deprecated
    public static short[] removeAllOccurences(final short[] array, final short element) {
        return (short[]) removeAt(array, indexesOf(array, element));
    }
    @Deprecated
    public static <T> T[] removeAllOccurences(final T[] array, final T element) {
        return (T[]) removeAt(array, indexesOf(array, element));
    }
    public static boolean[] removeAllOccurrences(final boolean[] array, final boolean element) {
        return (boolean[]) removeAt(array, indexesOf(array, element));
    }
    public static byte[] removeAllOccurrences(final byte[] array, final byte element) {
        return (byte[]) removeAt(array, indexesOf(array, element));
    }
    public static char[] removeAllOccurrences(final char[] array, final char element) {
        return (char[]) removeAt(array, indexesOf(array, element));
    }
    public static double[] removeAllOccurrences(final double[] array, final double element) {
        return (double[]) removeAt(array, indexesOf(array, element));
    }
    public static float[] removeAllOccurrences(final float[] array, final float element) {
        return (float[]) removeAt(array, indexesOf(array, element));
    }
    public static int[] removeAllOccurrences(final int[] array, final int element) {
        return (int[]) removeAt(array, indexesOf(array, element));
    }
    public static long[] removeAllOccurrences(final long[] array, final long element) {
        return (long[]) removeAt(array, indexesOf(array, element));
    }
    public static short[] removeAllOccurrences(final short[] array, final short element) {
        return (short[]) removeAt(array, indexesOf(array, element));
    }
    public static <T> T[] removeAllOccurrences(final T[] array, final T element) {
        return (T[]) removeAt(array, indexesOf(array, element));
    }
    static Object removeAt(final Object array, final BitSet indices) {
        if (array == null) {
            return null;
        }
        final int srcLength = getLength(array);
        final int removals = indices.cardinality(); 
        final Object result = Array.newInstance(array.getClass().getComponentType(), srcLength - removals);
        int srcIndex = 0;
        int destIndex = 0;
        int count;
        int set;
        while ((set = indices.nextSetBit(srcIndex)) != -1) {
            count = set - srcIndex;
            if (count > 0) {
                System.arraycopy(array, srcIndex, result, destIndex, count);
                destIndex += count;
            }
            srcIndex = indices.nextClearBit(set);
        }
        count = srcLength - srcIndex;
        if (count > 0) {
            System.arraycopy(array, srcIndex, result, destIndex, count);
        }
        return result;
    }
    public static boolean[] removeElement(final boolean[] array, final boolean element) {
        final int index = indexOf(array, element);
        return index == INDEX_NOT_FOUND ? clone(array) : remove(array, index);
    }
    public static byte[] removeElement(final byte[] array, final byte element) {
        final int index = indexOf(array, element);
        return index == INDEX_NOT_FOUND ? clone(array) : remove(array, index);
    }
    public static char[] removeElement(final char[] array, final char element) {
        final int index = indexOf(array, element);
        return index == INDEX_NOT_FOUND ? clone(array) : remove(array, index);
    }
    public static double[] removeElement(final double[] array, final double element) {
        final int index = indexOf(array, element);
        return index == INDEX_NOT_FOUND ? clone(array) : remove(array, index);
    }
    public static float[] removeElement(final float[] array, final float element) {
        final int index = indexOf(array, element);
        return index == INDEX_NOT_FOUND ? clone(array) : remove(array, index);
    }
    public static int[] removeElement(final int[] array, final int element) {
        final int index = indexOf(array, element);
        return index == INDEX_NOT_FOUND ? clone(array) : remove(array, index);
    }
    public static long[] removeElement(final long[] array, final long element) {
        final int index = indexOf(array, element);
        return index == INDEX_NOT_FOUND ? clone(array) : remove(array, index);
    }
    public static short[] removeElement(final short[] array, final short element) {
        final int index = indexOf(array, element);
        return index == INDEX_NOT_FOUND ? clone(array) : remove(array, index);
    }
    public static <T> T[] removeElement(final T[] array, final Object element) {
        final int index = indexOf(array, element);
        return index == INDEX_NOT_FOUND ? clone(array) : remove(array, index);
    }
    public static boolean[] removeElements(final boolean[] array, final boolean... values) {
        if (isEmpty(array) || isEmpty(values)) {
            return clone(array);
        }
        final HashMap<Boolean, MutableInt> occurrences = new HashMap<>(2); 
        for (final boolean v : values) {
            increment(occurrences, Boolean.valueOf(v));
        }
        final BitSet toRemove = new BitSet();
        for (int i = 0; i < array.length; i++) {
            final boolean key = array[i];
            final MutableInt count = occurrences.get(key);
            if (count != null) {
                if (count.decrementAndGet() == 0) {
                    occurrences.remove(key);
                }
                toRemove.set(i);
            }
        }
        return (boolean[]) removeAt(array, toRemove);
    }
    public static byte[] removeElements(final byte[] array, final byte... values) {
        if (isEmpty(array) || isEmpty(values)) {
            return clone(array);
        }
        final HashMap<Byte, MutableInt> occurrences = new HashMap<>(values.length);
        for (final byte v : values) {
            increment(occurrences, Byte.valueOf(v));
        }
        final BitSet toRemove = new BitSet();
        for (int i = 0; i < array.length; i++) {
            final byte key = array[i];
            final MutableInt count = occurrences.get(key);
            if (count != null) {
                if (count.decrementAndGet() == 0) {
                    occurrences.remove(key);
                }
                toRemove.set(i);
            }
        }
        return (byte[]) removeAt(array, toRemove);
    }
    public static char[] removeElements(final char[] array, final char... values) {
        if (isEmpty(array) || isEmpty(values)) {
            return clone(array);
        }
        final HashMap<Character, MutableInt> occurrences = new HashMap<>(values.length);
        for (final char v : values) {
            increment(occurrences, Character.valueOf(v));
        }
        final BitSet toRemove = new BitSet();
        for (int i = 0; i < array.length; i++) {
            final char key = array[i];
            final MutableInt count = occurrences.get(key);
            if (count != null) {
                if (count.decrementAndGet() == 0) {
                    occurrences.remove(key);
                }
                toRemove.set(i);
            }
        }
        return (char[]) removeAt(array, toRemove);
    }
    public static double[] removeElements(final double[] array, final double... values) {
        if (isEmpty(array) || isEmpty(values)) {
            return clone(array);
        }
        final HashMap<Double, MutableInt> occurrences = new HashMap<>(values.length);
        for (final double v : values) {
            increment(occurrences, Double.valueOf(v));
        }
        final BitSet toRemove = new BitSet();
        for (int i = 0; i < array.length; i++) {
            final double key = array[i];
            final MutableInt count = occurrences.get(key);
            if (count != null) {
                if (count.decrementAndGet() == 0) {
                    occurrences.remove(key);
                }
                toRemove.set(i);
            }
        }
        return (double[]) removeAt(array, toRemove);
    }
    public static float[] removeElements(final float[] array, final float... values) {
        if (isEmpty(array) || isEmpty(values)) {
            return clone(array);
        }
        final HashMap<Float, MutableInt> occurrences = new HashMap<>(values.length);
        for (final float v : values) {
            increment(occurrences, Float.valueOf(v));
        }
        final BitSet toRemove = new BitSet();
        for (int i = 0; i < array.length; i++) {
            final float key = array[i];
            final MutableInt count = occurrences.get(key);
            if (count != null) {
                if (count.decrementAndGet() == 0) {
                    occurrences.remove(key);
                }
                toRemove.set(i);
            }
        }
        return (float[]) removeAt(array, toRemove);
    }
    public static int[] removeElements(final int[] array, final int... values) {
        if (isEmpty(array) || isEmpty(values)) {
            return clone(array);
        }
        final HashMap<Integer, MutableInt> occurrences = new HashMap<>(values.length);
        for (final int v : values) {
            increment(occurrences, Integer.valueOf(v));
        }
        final BitSet toRemove = new BitSet();
        for (int i = 0; i < array.length; i++) {
            final int key = array[i];
            final MutableInt count = occurrences.get(key);
            if (count != null) {
                if (count.decrementAndGet() == 0) {
                    occurrences.remove(key);
                }
                toRemove.set(i);
            }
        }
        return (int[]) removeAt(array, toRemove);
    }
    public static long[] removeElements(final long[] array, final long... values) {
        if (isEmpty(array) || isEmpty(values)) {
            return clone(array);
        }
        final HashMap<Long, MutableInt> occurrences = new HashMap<>(values.length);
        for (final long v : values) {
            increment(occurrences, Long.valueOf(v));
        }
        final BitSet toRemove = new BitSet();
        for (int i = 0; i < array.length; i++) {
            final long key = array[i];
            final MutableInt count = occurrences.get(key);
            if (count != null) {
                if (count.decrementAndGet() == 0) {
                    occurrences.remove(key);
                }
                toRemove.set(i);
            }
        }
        return (long[]) removeAt(array, toRemove);
    }
    public static short[] removeElements(final short[] array, final short... values) {
        if (isEmpty(array) || isEmpty(values)) {
            return clone(array);
        }
        final HashMap<Short, MutableInt> occurrences = new HashMap<>(values.length);
        for (final short v : values) {
            increment(occurrences, Short.valueOf(v));
        }
        final BitSet toRemove = new BitSet();
        for (int i = 0; i < array.length; i++) {
            final short key = array[i];
            final MutableInt count = occurrences.get(key);
            if (count != null) {
                if (count.decrementAndGet() == 0) {
                    occurrences.remove(key);
                }
                toRemove.set(i);
            }
        }
        return (short[]) removeAt(array, toRemove);
    }
    @SafeVarargs
    public static <T> T[] removeElements(final T[] array, final T... values) {
        if (isEmpty(array) || isEmpty(values)) {
            return clone(array);
        }
        final HashMap<T, MutableInt> occurrences = new HashMap<>(values.length);
        for (final T v : values) {
            increment(occurrences, v);
        }
        final BitSet toRemove = new BitSet();
        for (int i = 0; i < array.length; i++) {
            final T key = array[i];
            final MutableInt count = occurrences.get(key);
            if (count != null) {
                if (count.decrementAndGet() == 0) {
                    occurrences.remove(key);
                }
                toRemove.set(i);
            }
        }
        @SuppressWarnings("unchecked") 
        final T[] result = (T[]) removeAt(array, toRemove);
        return result;
    }
    public static void reverse(final boolean[] array) {
        if (array != null) {
            reverse(array, 0, array.length);
        }
    }
    public static void reverse(final boolean[] array, final int startIndexInclusive, final int endIndexExclusive) {
        if (array == null) {
            return;
        }
        int i = Math.max(startIndexInclusive, 0);
        int j = max0(Math.min(array.length, endIndexExclusive)) - 1;
        boolean tmp;
        while (j > i) {
            tmp = array[j];
            array[j] = array[i];
            array[i] = tmp;
            j--;
            i++;
        }
    }
    public static void reverse(final byte[] array) {
        if (array != null) {
            reverse(array, 0, array.length);
        }
    }
    public static void reverse(final byte[] array, final int startIndexInclusive, final int endIndexExclusive) {
        if (array == null) {
            return;
        }
        int i = Math.max(startIndexInclusive, 0);
        int j = max0(Math.min(array.length, endIndexExclusive)) - 1;
        byte tmp;
        while (j > i) {
            tmp = array[j];
            array[j] = array[i];
            array[i] = tmp;
            j--;
            i++;
        }
    }
    public static void reverse(final char[] array) {
        if (array != null) {
            reverse(array, 0, array.length);
        }
    }
    public static void reverse(final char[] array, final int startIndexInclusive, final int endIndexExclusive) {
        if (array == null) {
            return;
        }
        int i = Math.max(startIndexInclusive, 0);
        int j = max0(Math.min(array.length, endIndexExclusive)) - 1;
        char tmp;
        while (j > i) {
            tmp = array[j];
            array[j] = array[i];
            array[i] = tmp;
            j--;
            i++;
        }
    }
    public static void reverse(final double[] array) {
        if (array != null) {
            reverse(array, 0, array.length);
        }
    }
    public static void reverse(final double[] array, final int startIndexInclusive, final int endIndexExclusive) {
        if (array == null) {
            return;
        }
        int i = Math.max(startIndexInclusive, 0);
        int j = max0(Math.min(array.length, endIndexExclusive)) - 1;
        double tmp;
        while (j > i) {
            tmp = array[j];
            array[j] = array[i];
            array[i] = tmp;
            j--;
            i++;
        }
    }
    public static void reverse(final float[] array) {
        if (array != null) {
            reverse(array, 0, array.length);
        }
    }
    public static void reverse(final float[] array, final int startIndexInclusive, final int endIndexExclusive) {
        if (array == null) {
            return;
        }
        int i = Math.max(startIndexInclusive, 0);
        int j = max0(Math.min(array.length, endIndexExclusive)) - 1;
        float tmp;
        while (j > i) {
            tmp = array[j];
            array[j] = array[i];
            array[i] = tmp;
            j--;
            i++;
        }
    }
    public static void reverse(final int[] array) {
        if (array != null) {
            reverse(array, 0, array.length);
        }
    }
    public static void reverse(final int[] array, final int startIndexInclusive, final int endIndexExclusive) {
        if (array == null) {
            return;
        }
        int i = Math.max(startIndexInclusive, 0);
        int j = max0(Math.min(array.length, endIndexExclusive)) - 1;
        int tmp;
        while (j > i) {
            tmp = array[j];
            array[j] = array[i];
            array[i] = tmp;
            j--;
            i++;
        }
    }
    public static void reverse(final long[] array) {
        if (array != null) {
            reverse(array, 0, array.length);
        }
    }
    public static void reverse(final long[] array, final int startIndexInclusive, final int endIndexExclusive) {
        if (array == null) {
            return;
        }
        int i = Math.max(startIndexInclusive, 0);
        int j = max0(Math.min(array.length, endIndexExclusive)) - 1;
        long tmp;
        while (j > i) {
            tmp = array[j];
            array[j] = array[i];
            array[i] = tmp;
            j--;
            i++;
        }
    }
    public static void reverse(final Object[] array) {
        if (array != null) {
            reverse(array, 0, array.length);
        }
    }
    public static void reverse(final Object[] array, final int startIndexInclusive, final int endIndexExclusive) {
        if (array == null) {
            return;
        }
        int i = Math.max(startIndexInclusive, 0);
        int j = max0(Math.min(array.length, endIndexExclusive)) - 1;
        Object tmp;
        while (j > i) {
            tmp = array[j];
            array[j] = array[i];
            array[i] = tmp;
            j--;
            i++;
        }
    }
    public static void reverse(final short[] array) {
        if (array != null) {
            reverse(array, 0, array.length);
        }
    }
    public static void reverse(final short[] array, final int startIndexInclusive, final int endIndexExclusive) {
        if (array == null) {
            return;
        }
        int i = Math.max(startIndexInclusive, 0);
        int j = max0(Math.min(array.length, endIndexExclusive)) - 1;
        short tmp;
        while (j > i) {
            tmp = array[j];
            array[j] = array[i];
            array[i] = tmp;
            j--;
            i++;
        }
    }
    public static <T> T[] setAll(final T[] array, final IntFunction<? extends T> generator) {
        if (array != null && generator != null) {
            Arrays.setAll(array, generator);
        }
        return array;
    }
    public static <T> T[] setAll(final T[] array, final Supplier<? extends T> generator) {
        if (array != null && generator != null) {
            for (int i = 0; i < array.length; i++) {
                array[i] = generator.get();
            }
        }
        return array;
    }
    public static void shift(final boolean[] array, final int offset) {
        if (array != null) {
            shift(array, 0, array.length, offset);
        }
    }
    public static void shift(final boolean[] array, int startIndexInclusive, int endIndexExclusive, int offset) {
        if (array == null || startIndexInclusive >= array.length - 1 || endIndexExclusive <= 0) {
            return;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = Math.min(endIndexExclusive, array.length);
        int n = endIndexExclusive - startIndexInclusive;
        if (n <= 1) {
            return;
        }
        offset %= n;
        if (offset < 0) {
            offset += n;
        }
        while (n > 1 && offset > 0) {
            final int nOffset = n - offset;
            if (offset > nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + n - nOffset,  nOffset);
                n = offset;
                offset -= nOffset;
            } else if (offset < nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset,  offset);
                startIndexInclusive += offset;
                n = nOffset;
            } else {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset, offset);
                break;
            }
        }
    }
    public static void shift(final byte[] array, final int offset) {
        if (array != null) {
            shift(array, 0, array.length, offset);
        }
    }
    public static void shift(final byte[] array, int startIndexInclusive, int endIndexExclusive, int offset) {
        if (array == null || startIndexInclusive >= array.length - 1 || endIndexExclusive <= 0) {
            return;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = Math.min(endIndexExclusive, array.length);
        int n = endIndexExclusive - startIndexInclusive;
        if (n <= 1) {
            return;
        }
        offset %= n;
        if (offset < 0) {
            offset += n;
        }
        while (n > 1 && offset > 0) {
            final int nOffset = n - offset;
            if (offset > nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + n - nOffset,  nOffset);
                n = offset;
                offset -= nOffset;
            } else if (offset < nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset,  offset);
                startIndexInclusive += offset;
                n = nOffset;
            } else {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset, offset);
                break;
            }
        }
    }
    public static void shift(final char[] array, final int offset) {
        if (array != null) {
            shift(array, 0, array.length, offset);
        }
    }
    public static void shift(final char[] array, int startIndexInclusive, int endIndexExclusive, int offset) {
        if (array == null || startIndexInclusive >= array.length - 1 || endIndexExclusive <= 0) {
            return;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = Math.min(endIndexExclusive, array.length);
        int n = endIndexExclusive - startIndexInclusive;
        if (n <= 1) {
            return;
        }
        offset %= n;
        if (offset < 0) {
            offset += n;
        }
        while (n > 1 && offset > 0) {
            final int nOffset = n - offset;
            if (offset > nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + n - nOffset,  nOffset);
                n = offset;
                offset -= nOffset;
            } else if (offset < nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset,  offset);
                startIndexInclusive += offset;
                n = nOffset;
            } else {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset, offset);
                break;
            }
        }
    }
    public static void shift(final double[] array, final int offset) {
        if (array != null) {
            shift(array, 0, array.length, offset);
        }
    }
    public static void shift(final double[] array, int startIndexInclusive, int endIndexExclusive, int offset) {
        if (array == null || startIndexInclusive >= array.length - 1 || endIndexExclusive <= 0) {
            return;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = Math.min(endIndexExclusive, array.length);
        int n = endIndexExclusive - startIndexInclusive;
        if (n <= 1) {
            return;
        }
        offset %= n;
        if (offset < 0) {
            offset += n;
        }
        while (n > 1 && offset > 0) {
            final int nOffset = n - offset;
            if (offset > nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + n - nOffset,  nOffset);
                n = offset;
                offset -= nOffset;
            } else if (offset < nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset,  offset);
                startIndexInclusive += offset;
                n = nOffset;
            } else {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset, offset);
                break;
            }
        }
    }
    public static void shift(final float[] array, final int offset) {
        if (array != null) {
            shift(array, 0, array.length, offset);
        }
    }
    public static void shift(final float[] array, int startIndexInclusive, int endIndexExclusive, int offset) {
        if (array == null || startIndexInclusive >= array.length - 1 || endIndexExclusive <= 0) {
            return;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = Math.min(endIndexExclusive, array.length);
        int n = endIndexExclusive - startIndexInclusive;
        if (n <= 1) {
            return;
        }
        offset %= n;
        if (offset < 0) {
            offset += n;
        }
        while (n > 1 && offset > 0) {
            final int nOffset = n - offset;
            if (offset > nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + n - nOffset,  nOffset);
                n = offset;
                offset -= nOffset;
            } else if (offset < nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset,  offset);
                startIndexInclusive += offset;
                n = nOffset;
            } else {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset, offset);
                break;
            }
        }
    }
    public static void shift(final int[] array, final int offset) {
        if (array != null) {
            shift(array, 0, array.length, offset);
        }
    }
    public static void shift(final int[] array, int startIndexInclusive, int endIndexExclusive, int offset) {
        if (array == null || startIndexInclusive >= array.length - 1 || endIndexExclusive <= 0) {
            return;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = Math.min(endIndexExclusive, array.length);
        int n = endIndexExclusive - startIndexInclusive;
        if (n <= 1) {
            return;
        }
        offset %= n;
        if (offset < 0) {
            offset += n;
        }
        while (n > 1 && offset > 0) {
            final int nOffset = n - offset;
            if (offset > nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + n - nOffset,  nOffset);
                n = offset;
                offset -= nOffset;
            } else if (offset < nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset,  offset);
                startIndexInclusive += offset;
                n = nOffset;
            } else {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset, offset);
                break;
            }
        }
    }
    public static void shift(final long[] array, final int offset) {
        if (array != null) {
            shift(array, 0, array.length, offset);
        }
    }
    public static void shift(final long[] array, int startIndexInclusive, int endIndexExclusive, int offset) {
        if (array == null || startIndexInclusive >= array.length - 1 || endIndexExclusive <= 0) {
            return;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = Math.min(endIndexExclusive, array.length);
        int n = endIndexExclusive - startIndexInclusive;
        if (n <= 1) {
            return;
        }
        offset %= n;
        if (offset < 0) {
            offset += n;
        }
        while (n > 1 && offset > 0) {
            final int nOffset = n - offset;
            if (offset > nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + n - nOffset,  nOffset);
                n = offset;
                offset -= nOffset;
            } else if (offset < nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset,  offset);
                startIndexInclusive += offset;
                n = nOffset;
            } else {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset, offset);
                break;
            }
        }
    }
    public static void shift(final Object[] array, final int offset) {
        if (array != null) {
            shift(array, 0, array.length, offset);
        }
    }
    public static void shift(final Object[] array, int startIndexInclusive, int endIndexExclusive, int offset) {
        if (array == null || startIndexInclusive >= array.length - 1 || endIndexExclusive <= 0) {
            return;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = Math.min(endIndexExclusive, array.length);
        int n = endIndexExclusive - startIndexInclusive;
        if (n <= 1) {
            return;
        }
        offset %= n;
        if (offset < 0) {
            offset += n;
        }
        while (n > 1 && offset > 0) {
            final int nOffset = n - offset;
            if (offset > nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + n - nOffset,  nOffset);
                n = offset;
                offset -= nOffset;
            } else if (offset < nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset,  offset);
                startIndexInclusive += offset;
                n = nOffset;
            } else {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset, offset);
                break;
            }
        }
    }
    public static void shift(final short[] array, final int offset) {
        if (array != null) {
            shift(array, 0, array.length, offset);
        }
    }
    public static void shift(final short[] array, int startIndexInclusive, int endIndexExclusive, int offset) {
        if (array == null || startIndexInclusive >= array.length - 1 || endIndexExclusive <= 0) {
            return;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = Math.min(endIndexExclusive, array.length);
        int n = endIndexExclusive - startIndexInclusive;
        if (n <= 1) {
            return;
        }
        offset %= n;
        if (offset < 0) {
            offset += n;
        }
        while (n > 1 && offset > 0) {
            final int nOffset = n - offset;
            if (offset > nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + n - nOffset,  nOffset);
                n = offset;
                offset -= nOffset;
            } else if (offset < nOffset) {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset,  offset);
                startIndexInclusive += offset;
                n = nOffset;
            } else {
                swap(array, startIndexInclusive, startIndexInclusive + nOffset, offset);
                break;
            }
        }
    }
    public static void shuffle(final boolean[] array) {
        shuffle(array, random());
    }
    public static void shuffle(final boolean[] array, final Random random) {
        if (array != null && random != null) {
            for (int i = array.length; i > 1; i--) {
                swap(array, i - 1, random.nextInt(i), 1);
            }
        }
    }
    public static void shuffle(final byte[] array) {
        shuffle(array, random());
    }
    public static void shuffle(final byte[] array, final Random random) {
        if (array != null && random != null) {
            for (int i = array.length; i > 1; i--) {
                swap(array, i - 1, random.nextInt(i), 1);
            }
        }
    }
    public static void shuffle(final char[] array) {
        shuffle(array, random());
    }
    public static void shuffle(final char[] array, final Random random) {
        if (array != null && random != null) {
            for (int i = array.length; i > 1; i--) {
                swap(array, i - 1, random.nextInt(i), 1);
            }
        }
    }
    public static void shuffle(final double[] array) {
        shuffle(array, random());
    }
    public static void shuffle(final double[] array, final Random random) {
        if (array != null && random != null) {
            for (int i = array.length; i > 1; i--) {
                swap(array, i - 1, random.nextInt(i), 1);
            }
        }
    }
    public static void shuffle(final float[] array) {
        shuffle(array, random());
    }
    public static void shuffle(final float[] array, final Random random) {
        if (array != null && random != null) {
            for (int i = array.length; i > 1; i--) {
                swap(array, i - 1, random.nextInt(i), 1);
            }
        }
    }
    public static void shuffle(final int[] array) {
        shuffle(array, random());
    }
    public static void shuffle(final int[] array, final Random random) {
        if (array != null && random != null) {
            for (int i = array.length; i > 1; i--) {
                swap(array, i - 1, random.nextInt(i), 1);
            }
        }
    }
    public static void shuffle(final long[] array) {
        shuffle(array, random());
    }
    public static void shuffle(final long[] array, final Random random) {
        if (array != null && random != null) {
            for (int i = array.length; i > 1; i--) {
                swap(array, i - 1, random.nextInt(i), 1);
            }
        }
    }
    public static void shuffle(final Object[] array) {
        shuffle(array, random());
    }
    public static void shuffle(final Object[] array, final Random random) {
        if (array != null && random != null) {
            for (int i = array.length; i > 1; i--) {
                swap(array, i - 1, random.nextInt(i), 1);
            }
        }
    }
    public static void shuffle(final short[] array) {
        shuffle(array, random());
    }
    public static void shuffle(final short[] array, final Random random) {
        if (array != null && random != null) {
            for (int i = array.length; i > 1; i--) {
                swap(array, i - 1, random.nextInt(i), 1);
            }
        }
    }
    public static boolean startsWith(final byte[] data, final byte[] expected) {
        if (data == expected) {
            return true;
        }
        if (data == null || expected == null) {
            return false;
        }
        final int dataLen = data.length;
        if (expected.length > dataLen) {
            return false;
        }
        if (expected.length == dataLen) {
            return Arrays.equals(data, expected);
        }
        for (int i = 0; i < expected.length; i++) {
            if (data[i] != expected[i]) {
                return false;
            }
        }
        return true;
    }
    public static boolean[] subarray(final boolean[] array, int startIndexInclusive, int endIndexExclusive) {
        if (array == null) {
            return null;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = max0(Math.min(endIndexExclusive, array.length));
        final int newSize = endIndexExclusive - startIndexInclusive;
        if (newSize <= 0) {
            return EMPTY_BOOLEAN_ARRAY;
        }
        return arraycopy(array, startIndexInclusive, 0, newSize, boolean[]::new);
    }
    public static byte[] subarray(final byte[] array, int startIndexInclusive, int endIndexExclusive) {
        if (array == null) {
            return null;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = max0(Math.min(endIndexExclusive, array.length));
        final int newSize = endIndexExclusive - startIndexInclusive;
        if (newSize <= 0) {
            return EMPTY_BYTE_ARRAY;
        }
        return arraycopy(array, startIndexInclusive, 0, newSize, byte[]::new);
    }
    public static char[] subarray(final char[] array, int startIndexInclusive, int endIndexExclusive) {
        if (array == null) {
            return null;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = max0(Math.min(endIndexExclusive, array.length));
        final int newSize = endIndexExclusive - startIndexInclusive;
        if (newSize <= 0) {
            return EMPTY_CHAR_ARRAY;
        }
        return arraycopy(array, startIndexInclusive, 0, newSize, char[]::new);
    }
    public static double[] subarray(final double[] array, int startIndexInclusive, int endIndexExclusive) {
        if (array == null) {
            return null;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = max0(Math.min(endIndexExclusive, array.length));
        final int newSize = endIndexExclusive - startIndexInclusive;
        if (newSize <= 0) {
            return EMPTY_DOUBLE_ARRAY;
        }
        return arraycopy(array, startIndexInclusive, 0, newSize, double[]::new);
    }
    public static float[] subarray(final float[] array, int startIndexInclusive, int endIndexExclusive) {
        if (array == null) {
            return null;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = max0(Math.min(endIndexExclusive, array.length));
        final int newSize = endIndexExclusive - startIndexInclusive;
        if (newSize <= 0) {
            return EMPTY_FLOAT_ARRAY;
        }
        return arraycopy(array, startIndexInclusive, 0, newSize, float[]::new);
    }
    public static int[] subarray(final int[] array, int startIndexInclusive, int endIndexExclusive) {
        if (array == null) {
            return null;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = max0(Math.min(endIndexExclusive, array.length));
        final int newSize = endIndexExclusive - startIndexInclusive;
        if (newSize <= 0) {
            return EMPTY_INT_ARRAY;
        }
        return arraycopy(array, startIndexInclusive, 0, newSize, int[]::new);
    }
    public static long[] subarray(final long[] array, int startIndexInclusive, int endIndexExclusive) {
        if (array == null) {
            return null;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = max0(Math.min(endIndexExclusive, array.length));
        final int newSize = endIndexExclusive - startIndexInclusive;
        if (newSize <= 0) {
            return EMPTY_LONG_ARRAY;
        }
        return arraycopy(array, startIndexInclusive, 0, newSize, long[]::new);
    }
    public static short[] subarray(final short[] array, int startIndexInclusive, int endIndexExclusive) {
        if (array == null) {
            return null;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = max0(Math.min(endIndexExclusive, array.length));
        final int newSize = endIndexExclusive - startIndexInclusive;
        if (newSize <= 0) {
            return EMPTY_SHORT_ARRAY;
        }
        return arraycopy(array, startIndexInclusive, 0, newSize, short[]::new);
    }
    public static <T> T[] subarray(final T[] array, int startIndexInclusive, int endIndexExclusive) {
        if (array == null) {
            return null;
        }
        startIndexInclusive = max0(startIndexInclusive);
        endIndexExclusive = max0(Math.min(endIndexExclusive, array.length));
        final int newSize = endIndexExclusive - startIndexInclusive;
        final Class<T> type = getComponentType(array);
        if (newSize <= 0) {
            return newInstance(type, 0);
        }
        return arraycopy(array, startIndexInclusive, 0, newSize, () -> newInstance(type, newSize));
    }
    public static void swap(final boolean[] array, final int offset1, final int offset2) {
        swap(array, offset1, offset2, 1);
    }
    public static void swap(final boolean[] array, int offset1, int offset2, int len) {
        if (isEmpty(array) || offset1 >= array.length || offset2 >= array.length) {
            return;
        }
        offset1 = max0(offset1);
        offset2 = max0(offset2);
        len = Math.min(Math.min(len, array.length - offset1), array.length - offset2);
        for (int i = 0; i < len; i++, offset1++, offset2++) {
            final boolean aux = array[offset1];
            array[offset1] = array[offset2];
            array[offset2] = aux;
        }
    }
    public static void swap(final byte[] array, final int offset1, final int offset2) {
        swap(array, offset1, offset2, 1);
    }
    public static void swap(final byte[] array, int offset1, int offset2, int len) {
        if (isEmpty(array) || offset1 >= array.length || offset2 >= array.length) {
            return;
        }
        offset1 = max0(offset1);
        offset2 = max0(offset2);
        len = Math.min(Math.min(len, array.length - offset1), array.length - offset2);
        for (int i = 0; i < len; i++, offset1++, offset2++) {
            final byte aux = array[offset1];
            array[offset1] = array[offset2];
            array[offset2] = aux;
        }
    }
    public static void swap(final char[] array, final int offset1, final int offset2) {
        swap(array, offset1, offset2, 1);
    }
    public static void swap(final char[] array, int offset1, int offset2, int len) {
        if (isEmpty(array) || offset1 >= array.length || offset2 >= array.length) {
            return;
        }
        offset1 = max0(offset1);
        offset2 = max0(offset2);
        len = Math.min(Math.min(len, array.length - offset1), array.length - offset2);
        for (int i = 0; i < len; i++, offset1++, offset2++) {
            final char aux = array[offset1];
            array[offset1] = array[offset2];
            array[offset2] = aux;
        }
    }
    public static void swap(final double[] array, final int offset1, final int offset2) {
        swap(array, offset1, offset2, 1);
    }
    public static void swap(final double[] array,  int offset1, int offset2, int len) {
        if (isEmpty(array) || offset1 >= array.length || offset2 >= array.length) {
            return;
        }
        offset1 = max0(offset1);
        offset2 = max0(offset2);
        len = Math.min(Math.min(len, array.length - offset1), array.length - offset2);
        for (int i = 0; i < len; i++, offset1++, offset2++) {
            final double aux = array[offset1];
            array[offset1] = array[offset2];
            array[offset2] = aux;
        }
    }
    public static void swap(final float[] array, final int offset1, final int offset2) {
        swap(array, offset1, offset2, 1);
    }
    public static void swap(final float[] array, int offset1, int offset2, int len) {
        if (isEmpty(array) || offset1 >= array.length || offset2 >= array.length) {
            return;
        }
        offset1 = max0(offset1);
        offset2 = max0(offset2);
        len = Math.min(Math.min(len, array.length - offset1), array.length - offset2);
        for (int i = 0; i < len; i++, offset1++, offset2++) {
            final float aux = array[offset1];
            array[offset1] = array[offset2];
            array[offset2] = aux;
        }
    }
    public static void swap(final int[] array, final int offset1, final int offset2) {
        swap(array, offset1, offset2, 1);
    }
    public static void swap(final int[] array,  int offset1, int offset2, int len) {
        if (isEmpty(array) || offset1 >= array.length || offset2 >= array.length) {
            return;
        }
        offset1 = max0(offset1);
        offset2 = max0(offset2);
        len = Math.min(Math.min(len, array.length - offset1), array.length - offset2);
        for (int i = 0; i < len; i++, offset1++, offset2++) {
            final int aux = array[offset1];
            array[offset1] = array[offset2];
            array[offset2] = aux;
        }
    }
    public static void swap(final long[] array, final int offset1, final int offset2) {
        swap(array, offset1, offset2, 1);
    }
    public static void swap(final long[] array,  int offset1, int offset2, int len) {
        if (isEmpty(array) || offset1 >= array.length || offset2 >= array.length) {
            return;
        }
        offset1 = max0(offset1);
        offset2 = max0(offset2);
        len = Math.min(Math.min(len, array.length - offset1), array.length - offset2);
        for (int i = 0; i < len; i++, offset1++, offset2++) {
            final long aux = array[offset1];
            array[offset1] = array[offset2];
            array[offset2] = aux;
        }
    }
    public static void swap(final Object[] array, final int offset1, final int offset2) {
        swap(array, offset1, offset2, 1);
    }
    public static void swap(final Object[] array,  int offset1, int offset2, int len) {
        if (isEmpty(array) || offset1 >= array.length || offset2 >= array.length) {
            return;
        }
        offset1 = max0(offset1);
        offset2 = max0(offset2);
        len = Math.min(Math.min(len, array.length - offset1), array.length - offset2);
        for (int i = 0; i < len; i++, offset1++, offset2++) {
            final Object aux = array[offset1];
            array[offset1] = array[offset2];
            array[offset2] = aux;
        }
    }
    public static void swap(final short[] array, final int offset1, final int offset2) {
        swap(array, offset1, offset2, 1);
    }
    public static void swap(final short[] array, int offset1, int offset2, int len) {
        if (isEmpty(array) || offset1 >= array.length || offset2 >= array.length) {
            return;
        }
        offset1 = max0(offset1);
        offset2 = max0(offset2);
        if (offset1 == offset2) {
            return;
        }
        len = Math.min(Math.min(len, array.length - offset1), array.length - offset2);
        for (int i = 0; i < len; i++, offset1++, offset2++) {
            final short aux = array[offset1];
            array[offset1] = array[offset2];
            array[offset2] = aux;
        }
    }
    public static <T> T[] toArray(@SuppressWarnings("unchecked") final T... items) {
        return items;
    }
    public static Map<Object, Object> toMap(final Object[] array) {
        if (array == null) {
            return null;
        }
        final Map<Object, Object> map = new HashMap<>((int) (array.length * 1.5));
        for (int i = 0; i < array.length; i++) {
            final Object object = array[i];
            if (object instanceof Map.Entry<?, ?>) {
                final Map.Entry<?, ?> entry = (Map.Entry<?, ?>) object;
                map.put(entry.getKey(), entry.getValue());
            } else if (object instanceof Object[]) {
                final Object[] entry = (Object[]) object;
                if (entry.length < 2) {
                    throw new IllegalArgumentException("Array element " + i + ", '"
                        + object
                        + "', has a length less than 2");
                }
                map.put(entry[0], entry[1]);
            } else {
                throw new IllegalArgumentException("Array element " + i + ", '"
                        + object
                        + "', is neither of type Map.Entry nor an Array");
            }
        }
        return map;
    }
    public static Boolean[] toObject(final boolean[] array) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_BOOLEAN_OBJECT_ARRAY;
        }
        return setAll(new Boolean[array.length], i -> array[i] ? Boolean.TRUE : Boolean.FALSE);
    }
    public static Byte[] toObject(final byte[] array) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_BYTE_OBJECT_ARRAY;
        }
        return setAll(new Byte[array.length], i -> Byte.valueOf(array[i]));
    }
    public static Character[] toObject(final char[] array) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_CHARACTER_OBJECT_ARRAY;
        }
        return setAll(new Character[array.length], i -> Character.valueOf(array[i]));
     }
    public static Double[] toObject(final double[] array) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_DOUBLE_OBJECT_ARRAY;
        }
        return setAll(new Double[array.length], i -> Double.valueOf(array[i]));
    }
    public static Float[] toObject(final float[] array) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_FLOAT_OBJECT_ARRAY;
        }
        return setAll(new Float[array.length], i -> Float.valueOf(array[i]));
    }
    public static Integer[] toObject(final int[] array) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_INTEGER_OBJECT_ARRAY;
        }
        return setAll(new Integer[array.length], i -> Integer.valueOf(array[i]));
    }
    public static Long[] toObject(final long[] array) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_LONG_OBJECT_ARRAY;
        }
        return setAll(new Long[array.length], i -> Long.valueOf(array[i]));
    }
    public static Short[] toObject(final short[] array) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_SHORT_OBJECT_ARRAY;
        }
        return setAll(new Short[array.length], i -> Short.valueOf(array[i]));
    }
    public static boolean[] toPrimitive(final Boolean[] array) {
        return toPrimitive(array, false);
    }
    public static boolean[] toPrimitive(final Boolean[] array, final boolean valueForNull) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_BOOLEAN_ARRAY;
        }
        final boolean[] result = new boolean[array.length];
        for (int i = 0; i < array.length; i++) {
            final Boolean b = array[i];
            result[i] = b == null ? valueForNull : b.booleanValue();
        }
        return result;
    }
    public static byte[] toPrimitive(final Byte[] array) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_BYTE_ARRAY;
        }
        final byte[] result = new byte[array.length];
        for (int i = 0; i < array.length; i++) {
            result[i] = array[i].byteValue();
        }
        return result;
    }
    public static byte[] toPrimitive(final Byte[] array, final byte valueForNull) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_BYTE_ARRAY;
        }
        final byte[] result = new byte[array.length];
        for (int i = 0; i < array.length; i++) {
            final Byte b = array[i];
            result[i] = b == null ? valueForNull : b.byteValue();
        }
        return result;
    }
    public static char[] toPrimitive(final Character[] array) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_CHAR_ARRAY;
        }
        final char[] result = new char[array.length];
        for (int i = 0; i < array.length; i++) {
            result[i] = array[i].charValue();
        }
        return result;
    }
    public static char[] toPrimitive(final Character[] array, final char valueForNull) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_CHAR_ARRAY;
        }
        final char[] result = new char[array.length];
        for (int i = 0; i < array.length; i++) {
            final Character b = array[i];
            result[i] = b == null ? valueForNull : b.charValue();
        }
        return result;
    }
    public static double[] toPrimitive(final Double[] array) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_DOUBLE_ARRAY;
        }
        final double[] result = new double[array.length];
        for (int i = 0; i < array.length; i++) {
            result[i] = array[i].doubleValue();
        }
        return result;
    }
    public static double[] toPrimitive(final Double[] array, final double valueForNull) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_DOUBLE_ARRAY;
        }
        final double[] result = new double[array.length];
        for (int i = 0; i < array.length; i++) {
            final Double b = array[i];
            result[i] = b == null ? valueForNull : b.doubleValue();
        }
        return result;
    }
    public static float[] toPrimitive(final Float[] array) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_FLOAT_ARRAY;
        }
        final float[] result = new float[array.length];
        for (int i = 0; i < array.length; i++) {
            result[i] = array[i].floatValue();
        }
        return result;
    }
    public static float[] toPrimitive(final Float[] array, final float valueForNull) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_FLOAT_ARRAY;
        }
        final float[] result = new float[array.length];
        for (int i = 0; i < array.length; i++) {
            final Float b = array[i];
            result[i] = b == null ? valueForNull : b.floatValue();
        }
        return result;
    }
    public static int[] toPrimitive(final Integer[] array) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_INT_ARRAY;
        }
        final int[] result = new int[array.length];
        for (int i = 0; i < array.length; i++) {
            result[i] = array[i].intValue();
        }
        return result;
    }
    public static int[] toPrimitive(final Integer[] array, final int valueForNull) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_INT_ARRAY;
        }
        final int[] result = new int[array.length];
        for (int i = 0; i < array.length; i++) {
            final Integer b = array[i];
            result[i] = b == null ? valueForNull : b.intValue();
        }
        return result;
    }
    public static long[] toPrimitive(final Long[] array) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_LONG_ARRAY;
        }
        final long[] result = new long[array.length];
        for (int i = 0; i < array.length; i++) {
            result[i] = array[i].longValue();
        }
        return result;
    }
    public static long[] toPrimitive(final Long[] array, final long valueForNull) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_LONG_ARRAY;
        }
        final long[] result = new long[array.length];
        for (int i = 0; i < array.length; i++) {
            final Long b = array[i];
            result[i] = b == null ? valueForNull : b.longValue();
        }
        return result;
    }
    public static Object toPrimitive(final Object array) {
        if (array == null) {
            return null;
        }
        final Class<?> ct = array.getClass().getComponentType();
        final Class<?> pt = ClassUtils.wrapperToPrimitive(ct);
        if (Boolean.TYPE.equals(pt)) {
            return toPrimitive((Boolean[]) array);
        }
        if (Character.TYPE.equals(pt)) {
            return toPrimitive((Character[]) array);
        }
        if (Byte.TYPE.equals(pt)) {
            return toPrimitive((Byte[]) array);
        }
        if (Integer.TYPE.equals(pt)) {
            return toPrimitive((Integer[]) array);
        }
        if (Long.TYPE.equals(pt)) {
            return toPrimitive((Long[]) array);
        }
        if (Short.TYPE.equals(pt)) {
            return toPrimitive((Short[]) array);
        }
        if (Double.TYPE.equals(pt)) {
            return toPrimitive((Double[]) array);
        }
        if (Float.TYPE.equals(pt)) {
            return toPrimitive((Float[]) array);
        }
        return array;
    }
    public static short[] toPrimitive(final Short[] array) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_SHORT_ARRAY;
        }
        final short[] result = new short[array.length];
        for (int i = 0; i < array.length; i++) {
            result[i] = array[i].shortValue();
        }
        return result;
    }
    public static short[] toPrimitive(final Short[] array, final short valueForNull) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_SHORT_ARRAY;
        }
        final short[] result = new short[array.length];
        for (int i = 0; i < array.length; i++) {
            final Short b = array[i];
            result[i] = b == null ? valueForNull : b.shortValue();
        }
        return result;
    }
    public static String toString(final Object array) {
        return toString(array, "{}");
    }
    public static String toString(final Object array, final String stringIfNull) {
        return array != null ? new ToStringBuilder(array, ToStringStyle.SIMPLE_STYLE).append(array).toString() : stringIfNull;
    }
    public static String[] toStringArray(final Object[] array) {
        return toStringArray(array, "null");
    }
    public static String[] toStringArray(final Object[] array, final String valueForNullElements) {
        if (array == null) {
            return null;
        }
        if (array.length == 0) {
            return EMPTY_STRING_ARRAY;
        }
        return map(array, String.class, e -> Objects.toString(e, valueForNullElements));
    }
    @Deprecated
    public ArrayUtils() {
    }
}
