package main

import(
    "github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/ebitenutil"
    "image/color"
)

type Sprite struct {
    image   *ebiten.Image
    x   float64
    y   float64
}

// Creating an image to add to the sprite
var square, _ = ebiten.NewImage(16, 16, ebiten.FilterNearest)

// Creating draw image options to add to the sprite
var opts = &ebiten.DrawImageOptions{}

// Creating the sprite instance
var sprite = Sprite{square, 0, 0,}


func update(screen *ebiten.Image) error {
    // Fill square
    sprite.image.Fill(color.White)

    opts := &ebiten.DrawImageOptions{}

    // When the "up" arrow key is pressed
    if ebiten.IsKeyPressed(ebiten.KeyUp){
        ebitenutil.DebugPrint(screen, "\nUp button")

        // Moving the sprite along the y axis
        sprite.y -= 1
    }

    // When the "down" arrow key is pressed
    if ebiten.IsKeyPressed(ebiten.KeyDown){
        ebitenutil.DebugPrint(screen, "\n\nDown button")

        // Moving the sprite along the y axis
        sprite.y += 1
    }

    if ebiten.IsKeyPressed(ebiten.KeyLeft){
        ebitenutil.DebugPrint(screen, "\n\n\nLeft key")

        // Moving the sprite along the x axis
        sprite.x -= 1
    }

    if ebiten.IsKeyPressed(ebiten.KeyRight){
        ebitenutil.DebugPrint(screen, "\n\n\n\nRight key")

        // Moving the sprite along the y axis
        sprite.x += 1
    }


    // Set the postion of sprite
    opts.GeoM.Translate(sprite.x, sprite.y)
    // Draw the square image on to the screen
    screen.DrawImage(sprite.image, opts)
    
    return nil
}

func main(){
    ebiten.Run(update, 320, 240, 2, "Hello image!")
}