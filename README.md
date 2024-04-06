
# ECSPOS Package Documentation

The `ecspos` package is designed to facilitate interactions with printers that support the ESC/POS protocol, providing a streamlined API for common printing tasks including text printing, alignment, and image processing.

## Types and Constructors

### `type Ecspos struct`

A struct that encapsulates the communication with an ESC/POS printer.

- `printer`: An `io.Writer` that represents the output stream to the printer.
- `content`: A slice of byte slices, each representing a command or data to send to the printer.
- `imageWidth`: An integer specifying the width to use when printing images.

### `func New(printer io.Writer) *Ecspos`

Creates a new `Ecspos` instance.

- **Parameters**:
  - `printer`: The output destination for print commands, typically an open file handle to a printer device.
- **Returns**: A pointer to an initialized `Ecspos` instance with default settings.

## Configuration Methods

### `func (e *Ecspos) SetImageWidth(w int) *Ecspos`

Sets the width for image printing operations. This can be useful for adjusting how images are scaled before being printed.

- **Parameters**:
  - `w`: The width in pixels.
- **Returns**: The `*Ecspos` instance to support method chaining.

### Text Alignment Methods

#### `func (e *Ecspos) AlignCenter() *Ecspos`

Configures the printer to align subsequent text centrally.

- **Returns**: The `*Ecspos` instance for chaining.

#### `func (e *Ecspos) AlignLeft() *Ecspos`

Sets text alignment to the left, which is the default state for most printers.

- **Returns**: The `*Ecspos` instance.

#### `func (e *Ecspos) AlignRight()`

Adjusts text alignment to the right.

## Font Style Methods

### `func (e *Ecspos) FontBold() *Ecspos`

Enables bold font style for printed text.

- **Returns**: The `*Ecspos` instance.

### `func (e *Ecspos) FontNormal() *Ecspos`

Resets the font style to normal (non-bold).

- **Returns**: The `*Ecspos` instance.

## Text Printing Methods

### `func (e *Ecspos) Text(txt string) *Ecspos`

Appends the provided text to the queue of content to be printed, followed by a newline character.

- **Parameters**:
  - `txt`: The text string to print.
- **Returns**: The `*Ecspos` instance.
