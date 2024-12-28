<script>
    export let isOpen = false;
    export let title = "Modal Title";
    export let message = "Modal Message";
    export let onClose; // Callback function for closing

    function close() {
        isOpen = false;
        if (onClose) {
            onClose();
        }
    }
</script>

{#if isOpen} 
    <div class="modal-overlay">
        <div class="modal">
            <h2>{title}</h2>
            <p>{message}</p>
            <button on:click={close}>Close</button>
        </div>
    </div>
{/if}



<style>
    .modal-overlay {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-color: rgba(0, 0, 0, 0.5); /* Semi-transparent background */
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .modal {
        background-color: white;
        padding: 20px;
        border-radius: 5px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
    }
</style>