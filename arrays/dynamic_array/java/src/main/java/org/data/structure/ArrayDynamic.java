package org.data.structure;

public class ArrayDynamic<T> {

    private Object[] array;
    private int length;     // number of elements in the array
    private int capacity;   // Will increase capacity x 2, every time we arrive to the limit

    public ArrayDynamic() {
        this.capacity = 1;
        this.array = new Object[this.capacity];
        this.length = 0;
    }

    // Return element at position index, if this position not defined will return null
    public T get(int index) {
        if (index < this.length && index >= 0)
            return (T) array[index];

        return null;
    }

    // Add element to the last position and return this value
    public T push(T item) {
        // Check if we have capacity
        if (this.length == this.capacity) {
            this.capacity *= 2;
            Object[] temp = new Object[this.capacity];

            for(int idx = 0; idx < this.length; idx++) {
                temp[idx] = this.array[idx];
            }

            this.array = temp;
        }

        this.array[this.length++] = item;

        return item;
    }

    // return and remove last element from array
    public T pop() {
        if (this.length > 0) {
            T item = (T) this.array[--this.length];
            this.array[this.length] = null;

            return item;
        }

        return null;
    }

    // Return element and position index and delete from array. If position not defined will return null
    public T deleteAtIndex(int index) {
        if (index < this.length && index >= 0) {
            T item = (T) this.array[index];
            shiftElements(index);

            return item;
        }

        return null;
    }

    // Delete element at position
    public void shiftElements(int index) {
        if (index < this.length && index >= 0) {
            for(int idx = index; idx < this.length - 1; idx++) {
                this.array[idx] = this.array[idx+1];
            }

            this.array[--this.length] = null;
        }
    }

    // Return size of the array
    public int getLength() {
        return this.length;
    }

    // Return capacity of the array
    public int getCapacity() {
        return this.capacity;
    }
}
